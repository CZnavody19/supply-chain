package db

import (
	"context"
	"errors"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

var ErrNotFound = errors.New("not found")

type DatabaseStore struct {
	db *neo4j.Driver
}

func NewDatabaseStore(db *neo4j.Driver) *DatabaseStore {
	return &DatabaseStore{
		db: db,
	}
}

func (ds *DatabaseStore) newSession(ctx context.Context) neo4j.SessionWithContext {
	return (*ds.db).NewSession(ctx, neo4j.SessionConfig{})
}

// ---- helpers to safely extract values from Neo4j records ----

func toString(val any) string {
	if val == nil {
		return ""
	}
	if s, ok := val.(string); ok {
		return s
	}
	return ""
}

func toFloat64(val any) float64 {
	if val == nil {
		return 0
	}
	switch v := val.(type) {
	case float64:
		return v
	case int64:
		return float64(v)
	}
	return 0
}

func toInt(val any) int {
	if val == nil {
		return 0
	}
	switch v := val.(type) {
	case int64:
		return int(v)
	case float64:
		return int(v)
	}
	return 0
}

func toSlice(val any) []any {
	if val == nil {
		return nil
	}
	if s, ok := val.([]any); ok {
		return s
	}
	return nil
}

func toMap(val any) map[string]any {
	if val == nil {
		return nil
	}
	if m, ok := val.(map[string]any); ok {
		return m
	}
	return nil
}
