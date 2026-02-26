package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/CZnavody19/supply-chain/src/db"
	"go.uber.org/zap"
)

// ---- Supply Chain Health ----

func (hh *HttpHandler) GetSupplyChainHealth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	health, err := hh.store.GetSupplyChainHealth(ctx)
	if err != nil {
		zap.L().Error("Error retrieving supply chain health", zap.Error(err))
		http.Error(w, "Error retrieving supply chain health", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(health)
}

// ---- Impact Analysis ----

func (hh *HttpHandler) GetImpactAnalysis(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	supplier := r.URL.Query().Get("supplier")
	if supplier == "" {
		http.Error(w, "Missing supplier parameter", http.StatusBadRequest)
		return
	}

	analysis, err := hh.store.GetImpactAnalysis(ctx, supplier)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			http.Error(w, "Supplier not found", http.StatusNotFound)
			return
		}
		zap.L().Error("Error retrieving impact analysis", zap.Error(err))
		http.Error(w, "Error retrieving impact analysis", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(analysis)
}

// ---- Forecast Delays ----

func (hh *HttpHandler) GetForecastDelays(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	monthsStr := r.URL.Query().Get("months")
	months := 3 // default
	if monthsStr != "" {
		m, err := strconv.Atoi(monthsStr)
		if err != nil {
			http.Error(w, "Invalid months parameter", http.StatusBadRequest)
			return
		}
		months = m
	}

	forecasts, err := hh.store.GetForecastDelays(ctx, months)
	if err != nil {
		zap.L().Error("Error retrieving delay forecasts", zap.Error(err))
		http.Error(w, "Error retrieving delay forecasts", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(forecasts)
}

// ---- Stock Levels Forecast ----

func (hh *HttpHandler) GetStockLevelForecast(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	product := r.URL.Query().Get("product")
	if product == "" {
		http.Error(w, "Missing product parameter", http.StatusBadRequest)
		return
	}

	monthsStr := r.URL.Query().Get("months")
	months := 6 // default
	if monthsStr != "" {
		m, err := strconv.Atoi(monthsStr)
		if err != nil {
			http.Error(w, "Invalid months parameter", http.StatusBadRequest)
			return
		}
		months = m
	}

	forecast, err := hh.store.GetStockLevelForecast(ctx, product, months)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		zap.L().Error("Error retrieving stock forecast", zap.Error(err))
		http.Error(w, "Error retrieving stock forecast", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(forecast)
}
