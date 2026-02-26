package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/CZnavody19/supply-chain/src/db"
	"go.uber.org/zap"
)

func (hh *HttpHandler) GetOptimalRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	if from == "" || to == "" {
		http.Error(w, "Missing 'from' or 'to' parameter", http.StatusBadRequest)
		return
	}

	result, err := hh.store.GetOptimalRoute(ctx, from, to)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			http.Error(w, "No route found between locations", http.StatusNotFound)
			return
		}
		zap.L().Error("Error finding optimal route", zap.Error(err))
		http.Error(w, "Error finding optimal route", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
