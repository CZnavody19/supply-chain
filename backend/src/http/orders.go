package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/CZnavody19/supply-chain/src/db"
	"github.com/CZnavody19/supply-chain/src/domain"
	"go.uber.org/zap"
)

func (hh *HttpHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")

	if id != "" {
		order, err := hh.store.GetOrderByID(ctx, id)
		if err != nil {
			if errors.Is(err, db.ErrNotFound) {
				http.Error(w, "Order not found", http.StatusNotFound)
				return
			}
			zap.L().Error("Error retrieving order", zap.Error(err))
			http.Error(w, "Error retrieving order", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(order)
		return
	}

	orders, err := hh.store.ListOrders(ctx)
	if err != nil {
		zap.L().Error("Error listing orders", zap.Error(err))
		http.Error(w, "Error listing orders", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}

type CreateOrderRequest struct {
	domain.Order
	ProductID  string  `json:"productId"`
	ProductQty int     `json:"productQuantity"`
	UnitPrice  float64 `json:"unitPrice"`
	CustomerID string  `json:"customerId"`
	SupplierID string  `json:"supplierId"`
}

func (hh *HttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		zap.L().Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := hh.store.CreateOrder(ctx, &req.Order, req.ProductID, req.ProductQty, req.UnitPrice, req.CustomerID, req.SupplierID); err != nil {
		zap.L().Error("Error creating order", zap.Error(err))
		http.Error(w, "Error creating order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req.Order)
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status"`
}

func (hh *HttpHandler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	var req UpdateOrderStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		zap.L().Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := hh.store.UpdateOrderStatus(ctx, id, req.Status); err != nil {
		zap.L().Error("Error updating order status", zap.Error(err))
		http.Error(w, "Error updating order status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// ---- Supply Path ----

func (hh *HttpHandler) GetSupplyPath(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	orderId := r.URL.Query().Get("orderId")
	if orderId == "" {
		http.Error(w, "Missing orderId parameter", http.StatusBadRequest)
		return
	}

	path, err := hh.store.GetSupplyPath(ctx, orderId)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}
		zap.L().Error("Error retrieving supply path", zap.Error(err))
		http.Error(w, "Error retrieving supply path", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(path)
}

// ---- Cost Breakdown ----

func (hh *HttpHandler) GetCostBreakdown(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	orderId := r.URL.Query().Get("orderId")
	if orderId == "" {
		http.Error(w, "Missing orderId parameter", http.StatusBadRequest)
		return
	}

	breakdown, err := hh.store.GetCostBreakdown(ctx, orderId)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}
		zap.L().Error("Error retrieving cost breakdown", zap.Error(err))
		http.Error(w, "Error retrieving cost breakdown", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(breakdown)
}
