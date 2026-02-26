package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/CZnavody19/supply-chain/src/db"
	"github.com/CZnavody19/supply-chain/src/domain"
	"go.uber.org/zap"
)

func (hh *HttpHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")

	if id != "" {
		product, err := hh.store.GetProductByID(ctx, id)
		if err != nil {
			if errors.Is(err, db.ErrNotFound) {
				http.Error(w, "Product not found", http.StatusNotFound)
				return
			}
			zap.L().Error("Error retrieving product", zap.Error(err))
			http.Error(w, "Error retrieving product", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(product)
		return
	}

	products, err := hh.store.ListProducts(ctx)
	if err != nil {
		zap.L().Error("Error listing products", zap.Error(err))
		http.Error(w, "Error listing products", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (hh *HttpHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var product domain.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		zap.L().Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := hh.store.CreateProduct(ctx, &product); err != nil {
		zap.L().Error("Error creating product", zap.Error(err))
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (hh *HttpHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	var product domain.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		zap.L().Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	product.ID = id

	if err := hh.store.UpdateProduct(ctx, &product); err != nil {
		zap.L().Error("Error updating product", zap.Error(err))
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (hh *HttpHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	if err := hh.store.DeleteProduct(ctx, id); err != nil {
		zap.L().Error("Error deleting product", zap.Error(err))
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ---- BOM ----

func (hh *HttpHandler) GetProductBOM(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	bom, err := hh.store.GetProductBOM(ctx, id)
	if err != nil {
		zap.L().Error("Error retrieving BOM", zap.Error(err))
		http.Error(w, "Error retrieving BOM", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bom)
}

func (hh *HttpHandler) GetProductBOMDetailed(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	bom, err := hh.store.GetProductBOMDetailed(ctx, id)
	if err != nil {
		zap.L().Error("Error retrieving detailed BOM", zap.Error(err))
		http.Error(w, "Error retrieving detailed BOM", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bom)
}

type AddBOMComponentRequest struct {
	ComponentID string `json:"componentId"`
	Quantity    int    `json:"quantity"`
	Position    int    `json:"position"`
}

func (hh *HttpHandler) AddBOMComponent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	var req AddBOMComponentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		zap.L().Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := hh.store.AddComponentToProduct(ctx, id, req.ComponentID, req.Quantity, req.Position); err != nil {
		zap.L().Error("Error adding component to product", zap.Error(err))
		http.Error(w, "Error adding component to product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type UpdateBOMComponentRequest struct {
	Quantity int `json:"quantity"`
}

func (hh *HttpHandler) UpdateBOMComponent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")
	componentID := r.URL.Query().Get("componentId")
	if id == "" || componentID == "" {
		http.Error(w, "Missing id or componentId parameter", http.StatusBadRequest)
		return
	}

	var req UpdateBOMComponentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		zap.L().Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := hh.store.UpdateBOMComponent(ctx, id, componentID, req.Quantity); err != nil {
		zap.L().Error("Error updating BOM component", zap.Error(err))
		http.Error(w, "Error updating BOM component", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// ---- Alternative Suppliers ----

func (hh *HttpHandler) GetAlternativeSuppliers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	suppliers, err := hh.store.GetAlternativeSuppliers(ctx, id)
	if err != nil {
		zap.L().Error("Error retrieving alternative suppliers", zap.Error(err))
		http.Error(w, "Error retrieving alternative suppliers", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(suppliers)
}
