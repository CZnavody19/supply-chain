package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/CZnavody19/supply-chain/src/db"
	"github.com/CZnavody19/supply-chain/src/domain"
	"go.uber.org/zap"
)

func (hh *HttpHandler) GetComponents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")

	if id != "" {
		component, err := hh.store.GetComponentByID(ctx, id)
		if err != nil {
			if errors.Is(err, db.ErrNotFound) {
				http.Error(w, "Component not found", http.StatusNotFound)
				return
			}
			zap.L().Error("Error retrieving component", zap.Error(err))
			http.Error(w, "Error retrieving component", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(component)
		return
	}

	components, err := hh.store.ListComponents(ctx)
	if err != nil {
		zap.L().Error("Error listing components", zap.Error(err))
		http.Error(w, "Error listing components", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(components)
}

func (hh *HttpHandler) CreateComponent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var component domain.Component
	if err := json.NewDecoder(r.Body).Decode(&component); err != nil {
		zap.L().Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := hh.store.CreateComponent(ctx, &component); err != nil {
		zap.L().Error("Error creating component", zap.Error(err))
		http.Error(w, "Error creating component", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(component)
}

func (hh *HttpHandler) UpdateComponent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	var component domain.Component
	if err := json.NewDecoder(r.Body).Decode(&component); err != nil {
		zap.L().Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	component.ID = id

	if err := hh.store.UpdateComponent(ctx, &component); err != nil {
		zap.L().Error("Error updating component", zap.Error(err))
		http.Error(w, "Error updating component", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(component)
}

func (hh *HttpHandler) DeleteComponent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	if err := hh.store.DeleteComponent(ctx, id); err != nil {
		zap.L().Error("Error deleting component", zap.Error(err))
		http.Error(w, "Error deleting component", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
