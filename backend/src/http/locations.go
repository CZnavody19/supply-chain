package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/CZnavody19/supply-chain/src/db"
	"github.com/CZnavody19/supply-chain/src/domain"
	"go.uber.org/zap"
)

func (hh *HttpHandler) GetLocations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")

	if id != "" {
		location, err := hh.store.GetLocationByID(ctx, id)
		if err != nil {
			if errors.Is(err, db.ErrNotFound) {
				http.Error(w, "Location not found", http.StatusNotFound)
				return
			}
			zap.L().Error("Error retrieving location", zap.Error(err))
			http.Error(w, "Error retrieving location", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(location)
		return
	}

	locations, err := hh.store.ListLocations(ctx)
	if err != nil {
		zap.L().Error("Error listing locations", zap.Error(err))
		http.Error(w, "Error listing locations", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(locations)
}

func (hh *HttpHandler) CreateLocation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var loc domain.Location
	if err := json.NewDecoder(r.Body).Decode(&loc); err != nil {
		zap.L().Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := hh.store.CreateLocation(ctx, &loc); err != nil {
		zap.L().Error("Error creating location", zap.Error(err))
		http.Error(w, "Error creating location", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(loc)
}

// ---- Inventory Status ----

func (hh *HttpHandler) GetInventoryStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	status, err := hh.store.GetInventoryStatus(ctx, id)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			http.Error(w, "Location not found", http.StatusNotFound)
			return
		}
		zap.L().Error("Error retrieving inventory status", zap.Error(err))
		http.Error(w, "Error retrieving inventory status", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(status)
}
