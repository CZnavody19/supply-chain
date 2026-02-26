package db

import (
	"context"

	"github.com/CZnavody19/supply-chain/src/domain"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func (ds *DatabaseStore) GetOptimalRoute(ctx context.Context, fromID, toID string) (*domain.OptimalRouteResult, error) {
	session := ds.newSession(ctx)
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			MATCH (from:Location {id: $from}), (to:Location {id: $to}),
			      path = shortestPath((from)-[:CONNECTED_TO*..10]-(to))
			UNWIND range(0, size(relationships(path))-1) AS idx
			WITH nodes(path) AS locs, relationships(path) AS rels, idx
			RETURN locs[idx].name AS fromName, locs[idx+1].name AS toName,
			       rels[idx].distance AS distance, rels[idx].time AS time, rels[idx].cost AS cost
		`, map[string]any{"from": fromID, "to": toID})
		if err != nil {
			return nil, err
		}

		var segments []domain.RouteSegment
		var totalDistance, totalTime, totalCost float64

		for res.Next(ctx) {
			rec := res.Record()
			fromV, _ := rec.Get("fromName")
			toV, _ := rec.Get("toName")
			distV, _ := rec.Get("distance")
			timeV, _ := rec.Get("time")
			costV, _ := rec.Get("cost")

			d := toFloat64(distV)
			t := toFloat64(timeV)
			c := toFloat64(costV)

			segments = append(segments, domain.RouteSegment{
				From:     toString(fromV),
				To:       toString(toV),
				Distance: d,
				Time:     t,
				Cost:     c,
			})
			totalDistance += d
			totalTime += t
			totalCost += c
		}
		if err = res.Err(); err != nil {
			return nil, err
		}

		if len(segments) == 0 {
			return nil, ErrNotFound
		}

		reliability := 1.0
		if len(segments) > 0 {
			reliability = max(0, 1.0-float64(len(segments))*0.02)
		}

		return &domain.OptimalRouteResult{
			Segments:         segments,
			TotalDistance:    totalDistance,
			TotalTime:        totalTime,
			TotalCost:        totalCost,
			TotalReliability: reliability,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*domain.OptimalRouteResult), nil
}
