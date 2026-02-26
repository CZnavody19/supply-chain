package db

import (
	"context"
	"fmt"
	"time"

	"github.com/CZnavody19/supply-chain/src/domain"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func (ds *DatabaseStore) ListOrders(ctx context.Context) ([]domain.Order, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (o:Order)
			RETURN o.id AS id, o.orderDate AS orderDate, o.dueDate AS dueDate,
			       o.quantity AS quantity, o.status AS status, o.cost AS cost
		`, nil)
		if err != nil {
			return nil, err
		}

		var orders []domain.Order
		for res.Next(ctx) {
			rec := res.Record()
			idV, _ := rec.Get("id")
			odV, _ := rec.Get("orderDate")
			ddV, _ := rec.Get("dueDate")
			qtyV, _ := rec.Get("quantity")
			statusV, _ := rec.Get("status")
			costV, _ := rec.Get("cost")

			orders = append(orders, domain.Order{
				ID:        toString(idV),
				OrderDate: toString(odV),
				DueDate:   toString(ddV),
				Quantity:  toInt(qtyV),
				Status:    toString(statusV),
				Cost:      toFloat64(costV),
			})
		}
		return orders, res.Err()
	})
	if err != nil {
		return nil, err
	}
	return result.([]domain.Order), nil
}

func (ds *DatabaseStore) GetOrderByID(ctx context.Context, id string) (*domain.Order, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (o:Order {id: $id})
			RETURN o.id AS id, o.orderDate AS orderDate, o.dueDate AS dueDate,
			       o.quantity AS quantity, o.status AS status, o.cost AS cost
		`, map[string]any{"id": id})
		if err != nil {
			return nil, err
		}

		if !res.Next(ctx) {
			return nil, ErrNotFound
		}
		rec := res.Record()
		idV, _ := rec.Get("id")
		odV, _ := rec.Get("orderDate")
		ddV, _ := rec.Get("dueDate")
		qtyV, _ := rec.Get("quantity")
		statusV, _ := rec.Get("status")
		costV, _ := rec.Get("cost")

		o := domain.Order{
			ID:        toString(idV),
			OrderDate: toString(odV),
			DueDate:   toString(ddV),
			Quantity:  toInt(qtyV),
			Status:    toString(statusV),
			Cost:      toFloat64(costV),
		}
		return &o, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*domain.Order), nil
}

func (ds *DatabaseStore) CreateOrder(ctx context.Context, order *domain.Order, productID string, productQty int, unitPrice float64, customerID, supplierID string) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	if order.ID == "" {
		order.ID = fmt.Sprintf("order-%d", time.Now().UnixNano())
	}

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
			CREATE (o:Order {
				id: $id, orderDate: $orderDate, dueDate: $dueDate,
				quantity: $quantity, status: $status, cost: $cost
			})
			WITH o
			MATCH (p:Product {id: $productId})
			CREATE (o)-[:CONTAINS {quantity: $prodQty, unitPrice: $unitPrice}]->(p)
			WITH o
			MATCH (cust:Company {id: $customerId})
			CREATE (o)-[:FROM]->(cust)
			WITH o
			MATCH (sup:Company {id: $supplierId})
			CREATE (o)-[:PLACED_WITH]->(sup)
		`, map[string]any{
			"id": order.ID, "orderDate": order.OrderDate, "dueDate": order.DueDate,
			"quantity": order.Quantity, "status": order.Status, "cost": order.Cost,
			"productId": productID, "prodQty": productQty, "unitPrice": unitPrice,
			"customerId": customerID, "supplierId": supplierID,
		})
		return nil, err
	})
	return err
}

func (ds *DatabaseStore) UpdateOrderStatus(ctx context.Context, id, status string) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
			MATCH (o:Order {id: $id})
			SET o.status = $status
		`, map[string]any{"id": id, "status": status})
		return nil, err
	})
	return err
}

// ---- Supply Path ----

func (ds *DatabaseStore) GetSupplyPath(ctx context.Context, orderID string) (*domain.SupplyPathResponse, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (o:Order {id: $id})-[cont:CONTAINS]->(p:Product)
			MATCH (o)-[:FROM]->(customer:Company)
			MATCH (o)-[:PLACED_WITH]->(supplier:Company)
			OPTIONAL MATCH (supplier)-[:LOCATED_AT]->(supLoc:Location)
			OPTIONAL MATCH (customer)-[:LOCATED_AT]->(custLoc:Location)
			OPTIONAL MATCH (o)-[sv:SHIPPED_VIA]->(r:Route)
			RETURN o.id AS orderId, o.quantity AS quantity, o.cost AS cost,
			       o.status AS orderStatus, o.dueDate AS dueDate,
			       p.name AS productName,
			       supplier.id AS supplierId, supplier.name AS supplierName, supplier.reliability AS supplierRel,
			       supLoc.id AS supLocId, supLoc.name AS supLocName,
			       customer.id AS customerId, customer.name AS customerName,
			       custLoc.id AS custLocId, custLoc.name AS custLocName,
			       r.distance AS routeDistance, r.estimatedTime AS routeTime, r.cost AS routeCost
		`, map[string]any{"id": orderID})
		if err != nil {
			return nil, err
		}

		if !res.Next(ctx) {
			return nil, ErrNotFound
		}
		rec := res.Record()

		prodNameV, _ := rec.Get("productName")
		qtyV, _ := rec.Get("quantity")
		costV, _ := rec.Get("cost")
		statusV, _ := rec.Get("orderStatus")
		dueDateV, _ := rec.Get("dueDate")
		supIdV, _ := rec.Get("supplierId")
		supNameV, _ := rec.Get("supplierName")
		supRelV, _ := rec.Get("supplierRel")
		supLocIdV, _ := rec.Get("supLocId")
		supLocNameV, _ := rec.Get("supLocName")
		custIdV, _ := rec.Get("customerId")
		custNameV, _ := rec.Get("customerName")
		custLocIdV, _ := rec.Get("custLocId")
		custLocNameV, _ := rec.Get("custLocName")
		routeDistV, _ := rec.Get("routeDistance")
		routeTimeV, _ := rec.Get("routeTime")
		routeCostV, _ := rec.Get("routeCost")

		var path []domain.SupplyPathStage

		// Stage 1: Source / Manufacturing
		path = append(path, domain.SupplyPathStage{
			Stage: 1,
			Name:  "Manufacturing",
			Company: &domain.CompanySummary{
				ID: toString(supIdV), Name: toString(supNameV), Reliability: toFloat64(supRelV),
			},
			Location: &domain.LocationSummary{
				ID: toString(supLocIdV), Name: toString(supLocNameV),
			},
			DueDate: toString(dueDateV),
			Status:  "completed",
		})

		// Stage 2: Transport (if route exists)
		if routeDistV != nil {
			path = append(path, domain.SupplyPathStage{
				Stage: 2,
				Name:  "Transport",
				From:  toString(supLocIdV),
				To:    toString(custLocIdV),
				Route: &domain.RouteSummary{
					Distance: toFloat64(routeDistV),
					Time:     fmt.Sprintf("%.0f hours", toFloat64(routeTimeV)),
					Cost:     toFloat64(routeCostV),
				},
				DueDate: toString(dueDateV),
				Status:  "in_transit",
			})
		}

		// Stage 3: Delivery
		path = append(path, domain.SupplyPathStage{
			Stage: len(path) + 1,
			Name:  "Delivery",
			Company: &domain.CompanySummary{
				ID: toString(custIdV), Name: toString(custNameV),
			},
			Location: &domain.LocationSummary{
				ID: toString(custLocIdV), Name: toString(custLocNameV),
			},
			DueDate: toString(dueDateV),
			Status:  toString(statusV),
		})

		var riskFactors []string
		if toFloat64(supRelV) < 0.95 {
			riskFactors = append(riskFactors, "Supplier reliability below 95%")
		}
		if routeDistV != nil && toFloat64(routeDistV) > 5000 {
			riskFactors = append(riskFactors, "Long distance route - weather delays possible")
		}

		totalDuration := "N/A"
		if routeTimeV != nil {
			totalDuration = fmt.Sprintf("%.0f hours", toFloat64(routeTimeV))
		}

		return &domain.SupplyPathResponse{
			OrderID:       orderID,
			Product:       toString(prodNameV),
			Quantity:      toInt(qtyV),
			TotalCost:     toFloat64(costV),
			Path:          path,
			TotalDuration: totalDuration,
			RiskFactors:   riskFactors,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*domain.SupplyPathResponse), nil
}

// ---- Cost Breakdown ----

func (ds *DatabaseStore) GetCostBreakdown(ctx context.Context, orderID string) (*domain.CostBreakdown, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (o:Order {id: $id})-[cont:CONTAINS]->(p:Product)
			OPTIONAL MATCH (p)-[bom:COMPOSED_OF]->(c:Component)
			OPTIONAL MATCH (o)-[:SHIPPED_VIA]->(r:Route)
			OPTIONAL MATCH (o)-[:PLACED_WITH]->(mfg:Company)-[mfgRel:MANUFACTURES]->(p)
			RETURN o.cost AS totalCost,
			       sum(c.price * bom.quantity) AS materialCost,
			       mfgRel.unitCost AS manufacturingUnitCost,
			       cont.quantity AS orderQty,
			       r.cost AS logisticsCost
		`, map[string]any{"id": orderID})
		if err != nil {
			return nil, err
		}

		if !res.Next(ctx) {
			return nil, ErrNotFound
		}
		rec := res.Record()
		totalCostV, _ := rec.Get("totalCost")
		matCostV, _ := rec.Get("materialCost")
		mfgUnitCostV, _ := rec.Get("manufacturingUnitCost")
		orderQtyV, _ := rec.Get("orderQty")
		logCostV, _ := rec.Get("logisticsCost")

		materialCost := toFloat64(matCostV)
		mfgCost := toFloat64(mfgUnitCostV) * float64(toInt(orderQtyV))
		logisticsCost := toFloat64(logCostV)
		totalCost := toFloat64(totalCostV)

		if totalCost == 0 {
			totalCost = materialCost + mfgCost + logisticsCost
		}

		return &domain.CostBreakdown{
			OrderID:           orderID,
			MaterialCost:      materialCost,
			ManufacturingCost: mfgCost,
			LogisticsCost:     logisticsCost,
			TotalCost:         totalCost,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*domain.CostBreakdown), nil
}
