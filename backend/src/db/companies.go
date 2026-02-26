package db

import (
	"context"
	"fmt"
	"time"

	"github.com/CZnavody19/supply-chain/src/domain"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func (ds *DatabaseStore) ListCompanies(ctx context.Context) ([]domain.Company, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (c:Company)
			RETURN c.id AS id, c.name AS name, c.type AS type, c.country AS country,
			       c.lat AS lat, c.lng AS lng, c.reliability AS reliability
		`, nil)
		if err != nil {
			return nil, err
		}

		var companies []domain.Company
		for res.Next(ctx) {
			rec := res.Record()
			idV, _ := rec.Get("id")
			nameV, _ := rec.Get("name")
			typeV, _ := rec.Get("type")
			countryV, _ := rec.Get("country")
			latV, _ := rec.Get("lat")
			lngV, _ := rec.Get("lng")
			relV, _ := rec.Get("reliability")

			companies = append(companies, domain.Company{
				ID:          toString(idV),
				Name:        toString(nameV),
				Type:        toString(typeV),
				Country:     toString(countryV),
				Coordinates: domain.Coordinates{Lat: toFloat64(latV), Lng: toFloat64(lngV)},
				Reliability: toFloat64(relV),
			})
		}
		return companies, res.Err()
	})
	if err != nil {
		return nil, err
	}
	return result.([]domain.Company), nil
}

func (ds *DatabaseStore) GetCompanyByID(ctx context.Context, id string) (*domain.Company, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (c:Company {id: $id})
			RETURN c.id AS id, c.name AS name, c.type AS type, c.country AS country,
			       c.lat AS lat, c.lng AS lng, c.reliability AS reliability
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
		countryV, _ := rec.Get("country")
		latV, _ := rec.Get("lat")
		lngV, _ := rec.Get("lng")
		relV, _ := rec.Get("reliability")

		c := domain.Company{
			ID:          toString(idV),
			Name:        toString(nameV),
			Type:        toString(typeV),
			Country:     toString(countryV),
			Coordinates: domain.Coordinates{Lat: toFloat64(latV), Lng: toFloat64(lngV)},
			Reliability: toFloat64(relV),
		}
		return &c, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*domain.Company), nil
}

func (ds *DatabaseStore) CreateCompany(ctx context.Context, company *domain.Company) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	if company.ID == "" {
		company.ID = fmt.Sprintf("c-%d", time.Now().UnixNano())
	}

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
			CREATE (c:Company {
				id: $id, name: $name, type: $type, country: $country,
				lat: $lat, lng: $lng, reliability: $reliability
			})
		`, map[string]any{
			"id": company.ID, "name": company.Name, "type": company.Type,
			"country": company.Country, "lat": company.Coordinates.Lat,
			"lng": company.Coordinates.Lng, "reliability": company.Reliability,
		})
		return nil, err
	})
	return err
}

func (ds *DatabaseStore) UpdateCompany(ctx context.Context, company *domain.Company) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `
			MATCH (c:Company {id: $id})
			SET c.name = $name, c.type = $type, c.country = $country,
			    c.lat = $lat, c.lng = $lng, c.reliability = $reliability
		`, map[string]any{
			"id": company.ID, "name": company.Name, "type": company.Type,
			"country": company.Country, "lat": company.Coordinates.Lat,
			"lng": company.Coordinates.Lng, "reliability": company.Reliability,
		})
		return nil, err
	})
	return err
}

func (ds *DatabaseStore) DeleteCompany(ctx context.Context, id string) error {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, `MATCH (c:Company {id: $id}) DETACH DELETE c`, map[string]any{"id": id})
		return nil, err
	})
	return err
}
