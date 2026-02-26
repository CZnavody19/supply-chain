package db

import (
	"context"
	"fmt"
	"time"

	"github.com/CZnavody19/supply-chain/src/domain"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

// ---- Risk Assessment ----

func (ds *DatabaseStore) GetRiskAssessment(ctx context.Context, companyID string) (*domain.RiskAssessment, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		// 1. Get company info
		compRes, err := tx.Run(ctx, `
			MATCH (c:Company {id: $id})
			RETURN c.name AS name, c.reliability AS reliability, c.country AS country
		`, map[string]any{"id": companyID})
		if err != nil {
			return nil, err
		}
		if !compRes.Next(ctx) {
			return nil, ErrNotFound
		}
		compRec := compRes.Record()
		compNameV, _ := compRec.Get("name")
		compRelV, _ := compRec.Get("reliability")

		// 2. Get order statistics
		statsRes, err := tx.Run(ctx, `
			MATCH (o:Order)-[:PLACED_WITH]->(c:Company {id: $id})
			RETURN count(o) AS total,
			       count(CASE WHEN o.status = 'delivered' THEN 1 END) AS delivered,
			       count(CASE WHEN o.status = 'delayed' THEN 1 END) AS delayed
		`, map[string]any{"id": companyID})
		if err != nil {
			return nil, err
		}

		var totalOrders, deliveredOrders, delayedOrders int
		if statsRes.Next(ctx) {
			rec := statsRes.Record()
			tV, _ := rec.Get("total")
			dV, _ := rec.Get("delivered")
			dlV, _ := rec.Get("delayed")
			totalOrders = toInt(tV)
			deliveredOrders = toInt(dV)
			delayedOrders = toInt(dlV)
		}

		// 3. Get critical products
		critRes, err := tx.Run(ctx, `
			MATCH (c:Company {id: $id})<-[:SUPPLIED_BY]-(comp:Component)<-[:COMPOSED_OF]-(p:Product)
			OPTIONAL MATCH (comp)-[:SUPPLIED_BY]->(alt:Company) WHERE alt.id <> $id
			WITH p, comp, count(DISTINCT alt) AS altCount
			RETURN p.name AS productName, comp.criticality AS criticality, altCount
		`, map[string]any{"id": companyID})
		if err != nil {
			return nil, err
		}

		var criticalProducts []domain.CriticalProduct
		for critRes.Next(ctx) {
			rec := critRes.Record()
			prodV, _ := rec.Get("productName")
			critV, _ := rec.Get("criticality")
			altV, _ := rec.Get("altCount")

			impact := "low"
			crit := toString(critV)
			if crit == "high" {
				impact = "high"
			} else if crit == "medium" {
				impact = "medium"
			}

			criticalProducts = append(criticalProducts, domain.CriticalProduct{
				Product:      toString(prodV),
				Impact:       impact,
				Alternatives: toInt(altV),
			})
		}

		// Compute scores
		reliability := toFloat64(compRelV)
		onTimeRate := 1.0
		if totalOrders > 0 {
			onTimeRate = float64(deliveredOrders) / float64(totalOrders)
		}
		qualityIssues := 0.0
		if totalOrders > 0 {
			qualityIssues = float64(delayedOrders) / float64(totalOrders)
		}
		riskScore := 1.0 - (reliability*0.4 + onTimeRate*0.4 + (1.0-qualityIssues)*0.2)

		// Generate recommendations
		var recommendations []string
		if reliability < 0.9 {
			recommendations = append(recommendations, "Supplier reliability is below 90%, consider diversifying")
		}
		if onTimeRate < 0.95 {
			recommendations = append(recommendations, "On-time delivery rate is below 95%, review logistics")
		}
		for _, cp := range criticalProducts {
			if cp.Alternatives < 2 && cp.Impact == "high" {
				recommendations = append(recommendations, fmt.Sprintf("Find alternative suppliers for critical product: %s", cp.Product))
			}
		}
		if len(recommendations) == 0 {
			recommendations = append(recommendations, "Supplier is performing well, continue monitoring")
		}

		return &domain.RiskAssessment{
			SupplierID: companyID,
			Company:    toString(compNameV),
			RiskScore:  riskScore,
			Factors: domain.RiskFactors{
				ReliabilityScore:   reliability,
				OnTimeDeliveryRate: onTimeRate,
				QualityIssues:      qualityIssues,
				GeopoliticalRisk:   0.5,
				FinancialStability: reliability,
			},
			CriticalFor:     criticalProducts,
			Recommendations: recommendations,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*domain.RiskAssessment), nil
}

// ---- Supply Chain Health ----

func (ds *DatabaseStore) GetSupplyChainHealth(ctx context.Context) (*domain.SupplyChainHealth, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		// 1. Critical components: high criticality with few suppliers
		critRes, err := tx.Run(ctx, `
			MATCH (c:Component)
			OPTIONAL MATCH (c)-[:SUPPLIED_BY]->(s:Company)
			WITH c, count(s) AS supplierCount
			WHERE c.criticality IN ['high', 'medium'] AND supplierCount <= 2
			RETURN c.id AS id, c.name AS name, c.criticality AS criticality, supplierCount
			ORDER BY supplierCount ASC
		`, nil)
		if err != nil {
			return nil, err
		}

		var criticalComponents []domain.CriticalComponentInfo
		for critRes.Next(ctx) {
			rec := critRes.Record()
			idV, _ := rec.Get("id")
			nameV, _ := rec.Get("name")
			critV, _ := rec.Get("criticality")
			scV, _ := rec.Get("supplierCount")
			criticalComponents = append(criticalComponents, domain.CriticalComponentInfo{
				ComponentID:   toString(idV),
				ComponentName: toString(nameV),
				Criticality:   toString(critV),
				SupplierCount: toInt(scV),
			})
		}

		// 2. Bottleneck locations: high stored quantity relative to capacity
		bottleRes, err := tx.Run(ctx, `
			MATCH (l:Location)
			WHERE l.capacity > 0
			OPTIONAL MATCH (p:Product)-[s:STORED_AT]->(l)
			WITH l, sum(coalesce(s.quantity, 0)) AS totalStored
			RETURN l.id AS id, l.name AS name,
			       toFloat(totalStored) / toFloat(l.capacity) AS utilization
			ORDER BY utilization DESC
			LIMIT 5
		`, nil)
		if err != nil {
			return nil, err
		}

		var bottlenecks []domain.BottleneckInfo
		for bottleRes.Next(ctx) {
			rec := bottleRes.Record()
			idV, _ := rec.Get("id")
			nameV, _ := rec.Get("name")
			utilV, _ := rec.Get("utilization")
			bottlenecks = append(bottlenecks, domain.BottleneckInfo{
				LocationID:   toString(idV),
				LocationName: toString(nameV),
				Utilization:  toFloat64(utilV),
			})
		}

		// 3. High-risk suppliers
		riskRes, err := tx.Run(ctx, `
			MATCH (c:Company)
			WHERE c.type = 'supplier' AND c.reliability < 0.9
			RETURN c.id AS id, c.name AS name, c.reliability AS reliability
			ORDER BY c.reliability ASC
		`, nil)
		if err != nil {
			return nil, err
		}

		var highRisk []domain.HighRiskSupplierInfo
		for riskRes.Next(ctx) {
			rec := riskRes.Record()
			idV, _ := rec.Get("id")
			nameV, _ := rec.Get("name")
			relV, _ := rec.Get("reliability")
			highRisk = append(highRisk, domain.HighRiskSupplierInfo{
				CompanyID:   toString(idV),
				CompanyName: toString(nameV),
				Reliability: toFloat64(relV),
			})
		}

		// Generate recommendations
		var recs []string
		if len(criticalComponents) > 0 {
			recs = append(recs, fmt.Sprintf("Found %d critical components with limited suppliers - diversify supply base", len(criticalComponents)))
		}
		for _, b := range bottlenecks {
			if b.Utilization > 0.8 {
				recs = append(recs, fmt.Sprintf("Location %s at %.0f%% capacity - consider expanding or redistributing", b.LocationName, b.Utilization*100))
			}
		}
		if len(highRisk) > 0 {
			recs = append(recs, fmt.Sprintf("%d high-risk suppliers detected - review contracts and find alternatives", len(highRisk)))
		}
		if len(recs) == 0 {
			recs = append(recs, "Supply chain is healthy - continue regular monitoring")
		}

		return &domain.SupplyChainHealth{
			CriticalComponents: criticalComponents,
			Bottlenecks:        bottlenecks,
			HighRiskSuppliers:  highRisk,
			Recommendations:    recs,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*domain.SupplyChainHealth), nil
}

// ---- Impact Analysis ----

func (ds *DatabaseStore) GetImpactAnalysis(ctx context.Context, supplierID string) (*domain.ImpactAnalysis, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		// 1. Get supplier name
		supRes, err := tx.Run(ctx, `
			MATCH (s:Company {id: $id})
			RETURN s.name AS name
		`, map[string]any{"id": supplierID})
		if err != nil {
			return nil, err
		}
		if !supRes.Next(ctx) {
			return nil, ErrNotFound
		}
		supRec := supRes.Record()
		supNameV, _ := supRec.Get("name")

		// 2. Get affected products and pending orders
		impactRes, err := tx.Run(ctx, `
			MATCH (s:Company {id: $id})<-[:SUPPLIED_BY]-(comp:Component)<-[:COMPOSED_OF]-(p:Product)
			OPTIONAL MATCH (o:Order)-[:CONTAINS]->(p) WHERE o.status IN ['pending', 'in_transit']
			WITH p, count(DISTINCT o) AS orderCount, sum(coalesce(o.cost, 0)) AS totalCost
			RETURN p.id AS productId, p.name AS productName, p.leadTime AS leadTime,
			       orderCount, totalCost
		`, map[string]any{"id": supplierID})
		if err != nil {
			return nil, err
		}

		var affected []domain.AffectedProduct
		var totalEstCost, totalRevenue float64
		for impactRes.Next(ctx) {
			rec := impactRes.Record()
			pidV, _ := rec.Get("productId")
			pnameV, _ := rec.Get("productName")
			ltV, _ := rec.Get("leadTime")
			ocV, _ := rec.Get("orderCount")
			tcV, _ := rec.Get("totalCost")

			cost := toFloat64(tcV)
			totalEstCost += cost
			totalRevenue += cost * 1.3 // rough margin estimate

			affected = append(affected, domain.AffectedProduct{
				ProductID:      toString(pidV),
				ProductName:    toString(pnameV),
				AffectedOrders: toInt(ocV),
				DelayDays:      toInt(ltV),
			})
		}

		// 3. Get alternative suppliers for mitigation
		altRes, err := tx.Run(ctx, `
			MATCH (s:Company {id: $id})<-[:SUPPLIED_BY]-(comp:Component)-[:SUPPLIED_BY]->(alt:Company)
			WHERE alt.id <> $id
			RETURN DISTINCT alt.name AS name, alt.reliability AS reliability
		`, map[string]any{"id": supplierID})
		if err != nil {
			return nil, err
		}

		var mitigation []string
		for altRes.Next(ctx) {
			rec := altRes.Record()
			nameV, _ := rec.Get("name")
			relV, _ := rec.Get("reliability")
			mitigation = append(mitigation, fmt.Sprintf("Switch to %s (reliability: %.0f%%)", toString(nameV), toFloat64(relV)*100))
		}
		mitigation = append(mitigation, "Use safety stock to cover transition period")
		mitigation = append(mitigation, "Review contracts with remaining suppliers for priority allocation")

		return &domain.ImpactAnalysis{
			SupplierID:   supplierID,
			SupplierName: toString(supNameV),
			Impact: domain.ImpactDetail{
				AffectedProducts: affected,
				EstimatedCost:    totalEstCost,
				AffectedRevenue:  totalRevenue,
				Mitigation:       mitigation,
			},
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*domain.ImpactAnalysis), nil
}

// ---- Forecast Delays ----

func (ds *DatabaseStore) GetForecastDelays(ctx context.Context, months int) ([]domain.ForecastDelay, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (o:Order)-[:CONTAINS]->(p:Product)
			WITH p,
			     count(o) AS totalOrders,
			     count(CASE WHEN o.status = 'delayed' THEN 1 END) AS delayedOrders
			WHERE totalOrders > 0
			RETURN p.id AS productId, p.name AS productName, p.leadTime AS leadTime,
			       totalOrders, delayedOrders,
			       toFloat(delayedOrders) / toFloat(totalOrders) AS delayRate
		`, nil)
		if err != nil {
			return nil, err
		}

		var forecasts []domain.ForecastDelay
		for res.Next(ctx) {
			rec := res.Record()
			pidV, _ := rec.Get("productId")
			pnameV, _ := rec.Get("productName")
			ltV, _ := rec.Get("leadTime")
			drV, _ := rec.Get("delayRate")

			delayRate := toFloat64(drV)
			leadTime := toFloat64(ltV)

			// Scale by requested time horizon
			probability := min(delayRate*float64(months)/3.0, 1.0)
			avgDelay := leadTime * delayRate

			riskLevel := "low"
			if probability > 0.5 {
				riskLevel = "high"
			} else if probability > 0.2 {
				riskLevel = "medium"
			}

			forecasts = append(forecasts, domain.ForecastDelay{
				ProductID:   toString(pidV),
				ProductName: toString(pnameV),
				AvgDelay:    avgDelay,
				Probability: probability,
				RiskLevel:   riskLevel,
			})
		}
		return forecasts, res.Err()
	})
	if err != nil {
		return nil, err
	}
	return result.([]domain.ForecastDelay), nil
}

// ---- Stock Level Forecast ----

func (ds *DatabaseStore) GetStockLevelForecast(ctx context.Context, productID string, months int) (*domain.StockLevelForecast, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		// 1. Get current total stock
		stockRes, err := tx.Run(ctx, `
			MATCH (p:Product {id: $id})
			OPTIONAL MATCH (p)-[s:STORED_AT]->(l:Location)
			RETURN p.name AS productName, sum(coalesce(s.quantity, 0)) AS totalStock
		`, map[string]any{"id": productID})
		if err != nil {
			return nil, err
		}
		if !stockRes.Next(ctx) {
			return nil, ErrNotFound
		}
		stockRec := stockRes.Record()
		pnameV, _ := stockRec.Get("productName")
		stockV, _ := stockRec.Get("totalStock")

		productName := toString(pnameV)
		currentStock := toInt(stockV)

		// 2. Get average consumption from delivered orders
		consumRes, err := tx.Run(ctx, `
			MATCH (o:Order)-[c:CONTAINS]->(p:Product {id: $id})
			WHERE o.status = 'delivered'
			RETURN sum(c.quantity) AS totalConsumed, count(o) AS orderCount
		`, map[string]any{"id": productID})
		if err != nil {
			return nil, err
		}

		avgMonthlyConsumption := 0.0
		if consumRes.Next(ctx) {
			rec := consumRes.Record()
			tcV, _ := rec.Get("totalConsumed")
			ocV, _ := rec.Get("orderCount")
			totalConsumed := toFloat64(tcV)
			orderCount := toFloat64(ocV)
			if orderCount > 0 {
				avgMonthlyConsumption = totalConsumed / max(orderCount, 1)
			}
		}

		// 3. Project stock levels per month
		now := time.Now()
		var projections []domain.StockProjection
		stock := float64(currentStock)
		for i := 1; i <= months; i++ {
			stock -= avgMonthlyConsumption
			projected := int(max(stock, 0))

			status := "sufficient"
			if projected <= 0 {
				status = "out_of_stock"
			} else if float64(projected) < avgMonthlyConsumption*2 {
				status = "low"
			}

			monthDate := now.AddDate(0, i, 0)
			projections = append(projections, domain.StockProjection{
				Month:          monthDate.Format("2006-01"),
				ProjectedStock: projected,
				Status:         status,
			})
		}

		return &domain.StockLevelForecast{
			ProductID:    productID,
			ProductName:  productName,
			CurrentStock: currentStock,
			Projections:  projections,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*domain.StockLevelForecast), nil
}
