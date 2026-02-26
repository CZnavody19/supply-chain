package db

import (
	"context"
	"fmt"
	"time"

	"github.com/CZnavody19/supply-chain/src/domain"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func (ds *DatabaseStore) ListLocations(ctx context.Context) ([]domain.Location, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (l:Location)
			RETURN l.id AS id, l.name AS name, l.type AS type,
			       l.lat AS lat, l.lng AS lng, l.capacity AS capacity
		`, nil)
		if err != nil {
			return nil, err
		}

		var locations []domain.Location
		for res.Next(ctx) {
			rec := res.Record()
			idV, _ := rec.Get("id")
			nameV, _ := rec.Get("name")
			typeV, _ := rec.Get("type")
			latV, _ := rec.Get("lat")
			lngV, _ := rec.Get("lng")
			capV, _ := rec.Get("capacity")

			locations = append(locations, domain.Location{
				ID:          toString(idV),
				Name:        toString(nameV),
				Type:        toString(typeV),
				Coordinates: domain.Coordinates{Lat: toFloat64(latV), Lng: toFloat64(lngV)},
				Capacity:    toInt(capV),
			})
		}
		return locations, res.Err()
	})
	if err != nil {
		return nil, err
	}
	return result.([]domain.Location), nil
}

func (ds *DatabaseStore) GetLocationByID(ctx context.Context, id string) (*domain.Location, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (l:Location {id: $id})
			RETURN l.id AS id, l.name AS name, l.type AS type,
			       l.lat AS lat, l.lng AS lng, l.capacity AS capacity
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
		typeV, _ := rec.Get("type")
		latV, _ := rec.Get("lat")
		lngV, _ := rec.Get("lng")
		capV, _ := rec.Get("capacity")

		l := domain.Location{
			ID:          toString(idV),
			Name:        toString(nameV),
			Type:        toString(typeV),
			Coordinates: domain.Coordinates{Lat: toFloat64(latV), Lng: toFloat64(lngV)},
			Capacity:    toInt(capV),
		}
		return &l, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*domain.Location), nil
}

func (ds *DatabaseStore) CreateLocation(ctx context.Context, loc *domain.Location) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	if loc.ID == "" {
		loc.ID = fmt.Sprintf("loc-%d", time.Now().UnixNano())
	}

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
			CREATE (l:Location {
				id: $id, name: $name, type: $type,
				lat: $lat, lng: $lng, capacity: $capacity
			})
		`, map[string]any{
			"id": loc.ID, "name": loc.Name, "type": loc.Type,
			"lat": loc.Coordinates.Lat, "lng": loc.Coordinates.Lng,
			"capacity": loc.Capacity,
		})
		return nil, err
	})
	return err
}

func (ds *DatabaseStore) GetInventoryStatus(ctx context.Context, locationID string) (*domain.InventoryStatus, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		// Get location info
		locRes, err := tx.Run(ctx, `
			MATCH (l:Location {id: $id})
			RETURN l.id AS id, l.name AS name, l.type AS type,
			       l.lat AS lat, l.lng AS lng, l.capacity AS capacity
		`, map[string]any{"id": locationID})
		if err != nil {
			return nil, err
		}
		if !locRes.Next(ctx) {
			return nil, ErrNotFound
		}
		locRec := locRes.Record()
		locIdV, _ := locRec.Get("id")
		locNameV, _ := locRec.Get("name")
		locTypeV, _ := locRec.Get("type")
		locLatV, _ := locRec.Get("lat")
		locLngV, _ := locRec.Get("lng")
		locCapV, _ := locRec.Get("capacity")

		location := domain.Location{
			ID:          toString(locIdV),
			Name:        toString(locNameV),
			Type:        toString(locTypeV),
			Coordinates: domain.Coordinates{Lat: toFloat64(locLatV), Lng: toFloat64(locLngV)},
			Capacity:    toInt(locCapV),
		}

		// Get stored products
		prodRes, err := tx.Run(ctx, `
			MATCH (p:Product)-[s:STORED_AT]->(l:Location {id: $id})
			OPTIONAL MATCH (o:Order)-[c:CONTAINS]->(p) WHERE o.status = 'delivered'
			WITH p, s, count(o) AS orderCount, sum(c.quantity) AS totalConsumed
			RETURN p.id AS id, p.name AS name, p.sku AS sku, p.price AS price,
			       p.weight AS weight, p.leadTime AS leadTime, p.status AS status,
			       s.quantity AS stockQty,
			       CASE WHEN orderCount > 0
			            THEN toInteger(toFloat(s.quantity) / (toFloat(totalConsumed) / toFloat(orderCount)))
			            ELSE 999 END AS daysOfSupply
		`, map[string]any{"id": locationID})
		if err != nil {
			return nil, err
		}

		var items []domain.InventoryItem
		for prodRes.Next(ctx) {
			rec := prodRes.Record()
			pIdV, _ := rec.Get("id")
			pNameV, _ := rec.Get("name")
			pSkuV, _ := rec.Get("sku")
			pPriceV, _ := rec.Get("price")
			pWeightV, _ := rec.Get("weight")
			pLtV, _ := rec.Get("leadTime")
			pStatusV, _ := rec.Get("status")
			stockV, _ := rec.Get("stockQty")
			dosV, _ := rec.Get("daysOfSupply")

			items = append(items, domain.InventoryItem{
				Product: domain.Product{
					ID:       toString(pIdV),
					Name:     toString(pNameV),
					SKU:      toString(pSkuV),
					Price:    toFloat64(pPriceV),
					Weight:   toFloat64(pWeightV),
					LeadTime: toInt(pLtV),
					Status:   toString(pStatusV),
				},
				Quantity:     toInt(stockV),
				DaysOfSupply: toInt(dosV),
			})
		}

		return &domain.InventoryStatus{
			Location: location,
			Products: items,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*domain.InventoryStatus), nil
}
