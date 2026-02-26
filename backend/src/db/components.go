package db

import (
	"context"
	"fmt"
	"time"

	"github.com/CZnavody19/supply-chain/src/domain"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func (ds *DatabaseStore) ListComponents(ctx context.Context) ([]domain.Component, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (c:Component)
			RETURN c.id AS id, c.name AS name, c.price AS price,
			       c.quantity AS quantity, c.criticality AS criticality
		`, nil)
		if err != nil {
			return nil, err
		}

		var components []domain.Component
		for res.Next(ctx) {
			rec := res.Record()
			idV, _ := rec.Get("id")
			nameV, _ := rec.Get("name")
			priceV, _ := rec.Get("price")
			qtyV, _ := rec.Get("quantity")
			critV, _ := rec.Get("criticality")

			components = append(components, domain.Component{
				ID:          toString(idV),
				Name:        toString(nameV),
				Price:       toFloat64(priceV),
				Quantity:    toInt(qtyV),
				Criticality: toString(critV),
			})
		}
		return components, res.Err()
	})
	if err != nil {
		return nil, err
	}
	return result.([]domain.Component), nil
}

func (ds *DatabaseStore) GetComponentByID(ctx context.Context, id string) (*domain.Component, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (c:Component {id: $id})
			RETURN c.id AS id, c.name AS name, c.price AS price,
			       c.quantity AS quantity, c.criticality AS criticality
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
		priceV, _ := rec.Get("price")
		qtyV, _ := rec.Get("quantity")
		critV, _ := rec.Get("criticality")

		c := domain.Component{
			ID:          toString(idV),
			Name:        toString(nameV),
			Price:       toFloat64(priceV),
			Quantity:    toInt(qtyV),
			Criticality: toString(critV),
		}
		return &c, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*domain.Component), nil
}

func (ds *DatabaseStore) CreateComponent(ctx context.Context, component *domain.Component) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	if component.ID == "" {
		component.ID = fmt.Sprintf("comp-%d", time.Now().UnixNano())
	}

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
			CREATE (c:Component {
				id: $id, name: $name, price: $price,
				quantity: $quantity, criticality: $criticality
			})
		`, map[string]any{
			"id": component.ID, "name": component.Name, "price": component.Price,
			"quantity": component.Quantity, "criticality": component.Criticality,
		})
		return nil, err
	})
	return err
}

func (ds *DatabaseStore) UpdateComponent(ctx context.Context, component *domain.Component) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
			MATCH (c:Component {id: $id})
			SET c.name = $name, c.price = $price,
			    c.quantity = $quantity, c.criticality = $criticality
		`, map[string]any{
			"id": component.ID, "name": component.Name, "price": component.Price,
			"quantity": component.Quantity, "criticality": component.Criticality,
		})
		return nil, err
	})
	return err
}

func (ds *DatabaseStore) DeleteComponent(ctx context.Context, id string) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `MATCH (c:Component {id: $id}) DETACH DELETE c`, map[string]any{"id": id})
		return nil, err
	})
	return err
}
