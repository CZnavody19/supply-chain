package db

import (
	"context"
	"fmt"
	"time"

	"github.com/CZnavody19/supply-chain/src/domain"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func (ds *DatabaseStore) ListProducts(ctx context.Context) ([]domain.Product, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (p:Product)
			RETURN p.id AS id, p.name AS name, p.sku AS sku, p.price AS price,
			       p.weight AS weight, p.leadTime AS leadTime, p.status AS status
		`, nil)
		if err != nil {
			return nil, err
		}

		var products []domain.Product
		for res.Next(ctx) {
			rec := res.Record()
			idV, _ := rec.Get("id")
			nameV, _ := rec.Get("name")
			skuV, _ := rec.Get("sku")
			priceV, _ := rec.Get("price")
			weightV, _ := rec.Get("weight")
			ltV, _ := rec.Get("leadTime")
			statusV, _ := rec.Get("status")

			products = append(products, domain.Product{
				ID:       toString(idV),
				Name:     toString(nameV),
				SKU:      toString(skuV),
				Price:    toFloat64(priceV),
				Weight:   toFloat64(weightV),
				LeadTime: toInt(ltV),
				Status:   toString(statusV),
			})
		}
		if err = res.Err(); err != nil {
			return nil, err
		}
		return products, nil
	})
	if err != nil {
		return nil, err
	}
	return result.([]domain.Product), nil
}

func (ds *DatabaseStore) GetProductByID(ctx context.Context, id string) (*domain.Product, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (p:Product {id: $id})
			RETURN p.id AS id, p.name AS name, p.sku AS sku, p.price AS price,
			       p.weight AS weight, p.leadTime AS leadTime, p.status AS status
		`, map[string]any{"id": id})
		if err != nil {
			return nil, err
		}

		if !res.Next(ctx) {
			return nil, ErrNotFound
		}
		rec := res.Record()
		idV, _ := rec.Get("id")
		nameV, _ := rec.Get("name")
		skuV, _ := rec.Get("sku")
		priceV, _ := rec.Get("price")
		weightV, _ := rec.Get("weight")
		ltV, _ := rec.Get("leadTime")
		statusV, _ := rec.Get("status")

		p := domain.Product{
			ID:       toString(idV),
			Name:     toString(nameV),
			SKU:      toString(skuV),
			Price:    toFloat64(priceV),
			Weight:   toFloat64(weightV),
			LeadTime: toInt(ltV),
			Status:   toString(statusV),
		}
		return &p, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*domain.Product), nil
}

func (ds *DatabaseStore) CreateProduct(ctx context.Context, product *domain.Product) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	if product.ID == "" {
		product.ID = fmt.Sprintf("prod-%d", time.Now().UnixNano())
	}

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
			CREATE (p:Product {
				id: $id, name: $name, sku: $sku, price: $price,
				weight: $weight, leadTime: $leadTime, status: $status
			})
		`, map[string]any{
			"id": product.ID, "name": product.Name, "sku": product.SKU,
			"price": product.Price, "weight": product.Weight,
			"leadTime": product.LeadTime, "status": product.Status,
		})
		return nil, err
	})
	return err
}

func (ds *DatabaseStore) UpdateProduct(ctx context.Context, product *domain.Product) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
			MATCH (p:Product {id: $id})
			SET p.name = $name, p.sku = $sku, p.price = $price,
			    p.weight = $weight, p.leadTime = $leadTime, p.status = $status
		`, map[string]any{
			"id": product.ID, "name": product.Name, "sku": product.SKU,
			"price": product.Price, "weight": product.Weight,
			"leadTime": product.LeadTime, "status": product.Status,
		})
		return nil, err
	})
	return err
}

func (ds *DatabaseStore) DeleteProduct(ctx context.Context, id string) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `MATCH (p:Product {id: $id}) DETACH DELETE p`, map[string]any{"id": id})
		return nil, err
	})
	return err
}

// ---- BOM (Bill of Materials) ----

func (ds *DatabaseStore) GetProductBOM(ctx context.Context, productID string) ([]domain.BOMEntry, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (p:Product {id: $id})-[r:COMPOSED_OF]->(c:Component)
			RETURN c.id AS id, c.name AS name, c.price AS price, c.criticality AS criticality,
			       r.quantity AS quantity, r.position AS position
			ORDER BY r.position
		`, map[string]any{"id": productID})
		if err != nil {
			return nil, err
		}

		var entries []domain.BOMEntry
		for res.Next(ctx) {
			rec := res.Record()
			idV, _ := rec.Get("id")
			nameV, _ := rec.Get("name")
			priceV, _ := rec.Get("price")
			critV, _ := rec.Get("criticality")
			qtyV, _ := rec.Get("quantity")
			posV, _ := rec.Get("position")

			entries = append(entries, domain.BOMEntry{
				Component: domain.Component{
					ID:          toString(idV),
					Name:        toString(nameV),
					Price:       toFloat64(priceV),
					Criticality: toString(critV),
				},
				Quantity: toInt(qtyV),
				Position: toInt(posV),
			})
		}
		return entries, res.Err()
	})
	if err != nil {
		return nil, err
	}
	return result.([]domain.BOMEntry), nil
}

func (ds *DatabaseStore) GetProductBOMDetailed(ctx context.Context, productID string) ([]domain.BOMDetailedEntry, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (p:Product {id: $id})-[r:COMPOSED_OF]->(c:Component)
			OPTIONAL MATCH (c)-[s:SUPPLIED_BY]->(sup:Company)
			WITH c, r, collect({
				companyId: sup.id, companyName: sup.name, companyType: sup.type,
				companyCountry: sup.country, companyReliability: sup.reliability,
				price: s.price, leadTime: s.leadTime, minOrder: s.minOrder
			}) AS suppliers
			RETURN c.id AS id, c.name AS name, c.price AS price, c.criticality AS criticality,
			       r.quantity AS quantity, r.position AS position, suppliers
			ORDER BY r.position
		`, map[string]any{"id": productID})
		if err != nil {
			return nil, err
		}

		var entries []domain.BOMDetailedEntry
		for res.Next(ctx) {
			rec := res.Record()
			idV, _ := rec.Get("id")
			nameV, _ := rec.Get("name")
			priceV, _ := rec.Get("price")
			critV, _ := rec.Get("criticality")
			qtyV, _ := rec.Get("quantity")
			posV, _ := rec.Get("position")
			suppV, _ := rec.Get("suppliers")

			var suppliers []domain.ComponentSupplier
			for _, s := range toSlice(suppV) {
				sm := toMap(s)
				if sm == nil || sm["companyId"] == nil {
					continue
				}
				suppliers = append(suppliers, domain.ComponentSupplier{
					Company: domain.Company{
						ID:          toString(sm["companyId"]),
						Name:        toString(sm["companyName"]),
						Type:        toString(sm["companyType"]),
						Country:     toString(sm["companyCountry"]),
						Reliability: toFloat64(sm["companyReliability"]),
					},
					Price:    toFloat64(sm["price"]),
					LeadTime: toInt(sm["leadTime"]),
					MinOrder: toInt(sm["minOrder"]),
				})
			}

			entries = append(entries, domain.BOMDetailedEntry{
				Component: domain.Component{
					ID:          toString(idV),
					Name:        toString(nameV),
					Price:       toFloat64(priceV),
					Criticality: toString(critV),
				},
				Quantity:  toInt(qtyV),
				Position:  toInt(posV),
				Suppliers: suppliers,
			})
		}
		return entries, res.Err()
	})
	if err != nil {
		return nil, err
	}
	return result.([]domain.BOMDetailedEntry), nil
}

func (ds *DatabaseStore) AddComponentToProduct(ctx context.Context, productID, componentID string, quantity, position int) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
			MATCH (p:Product {id: $productId}), (c:Component {id: $componentId})
			CREATE (p)-[:COMPOSED_OF {quantity: $quantity, position: $position}]->(c)
		`, map[string]any{
			"productId": productID, "componentId": componentID,
			"quantity": quantity, "position": position,
		})
		return nil, err
	})
	return err
}

func (ds *DatabaseStore) UpdateBOMComponent(ctx context.Context, productID, componentID string, quantity int) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
			MATCH (p:Product {id: $productId})-[r:COMPOSED_OF]->(c:Component {id: $componentId})
			SET r.quantity = $quantity
		`, map[string]any{
			"productId": productID, "componentId": componentID, "quantity": quantity,
		})
		return nil, err
	})
	return err
}

// ---- Alternative Suppliers ----

func (ds *DatabaseStore) GetAlternativeSuppliers(ctx context.Context, productID string) ([]domain.AlternativeSupplier, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (p:Product {id: $id})-[:COMPOSED_OF]->(c:Component)-[s:SUPPLIED_BY]->(sup:Company)
			RETURN DISTINCT sup.id AS id, sup.name AS name, sup.type AS type,
			       sup.country AS country, sup.reliability AS reliability,
			       s.price AS price, s.leadTime AS leadTime
			ORDER BY s.price ASC, sup.reliability DESC
		`, map[string]any{"id": productID})
		if err != nil {
			return nil, err
		}

		var suppliers []domain.AlternativeSupplier
		for res.Next(ctx) {
			rec := res.Record()
			idV, _ := rec.Get("id")
			nameV, _ := rec.Get("name")
			typeV, _ := rec.Get("type")
			countryV, _ := rec.Get("country")
			relV, _ := rec.Get("reliability")
			priceV, _ := rec.Get("price")
			ltV, _ := rec.Get("leadTime")

			suppliers = append(suppliers, domain.AlternativeSupplier{
				Company: domain.Company{
					ID:          toString(idV),
					Name:        toString(nameV),
					Type:        toString(typeV),
					Country:     toString(countryV),
					Reliability: toFloat64(relV),
				},
				Price:       toFloat64(priceV),
				Reliability: toFloat64(relV),
				LeadTime:    toInt(ltV),
			})
		}
		return suppliers, res.Err()
	})
	if err != nil {
		return nil, err
	}
	return result.([]domain.AlternativeSupplier), nil
}
