package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/CZnavody19/supply-chain/src/db"
	"github.com/CZnavody19/supply-chain/src/domain"
	"go.uber.org/zap"
)

func (hh *HttpHandler) GetCompanies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")

	if id != "" {
		company, err := hh.store.GetCompanyByID(ctx, id)
		if err != nil {
			if errors.Is(err, db.ErrNotFound) {
				http.Error(w, "Company not found", http.StatusNotFound)
				return
			}
			zap.L().Error("Error retrieving company", zap.Error(err))
			http.Error(w, "Error retrieving company", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(company)
		return
	}

	companies, err := hh.store.ListCompanies(ctx)
	if err != nil {
		zap.L().Error("Error listing companies", zap.Error(err))
		http.Error(w, "Error listing companies", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(companies)
}

func (hh *HttpHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var company domain.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		zap.L().Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := hh.store.CreateCompany(ctx, &company); err != nil {
		zap.L().Error("Error creating company", zap.Error(err))
		http.Error(w, "Error creating company", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(company)
}

func (hh *HttpHandler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	var company domain.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		zap.L().Error("Failed to decode request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	company.ID = id

	if err := hh.store.UpdateCompany(ctx, &company); err != nil {
		zap.L().Error("Error updating company", zap.Error(err))
		http.Error(w, "Error updating company", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(company)
}

func (hh *HttpHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	if err := hh.store.DeleteCompany(ctx, id); err != nil {
		zap.L().Error("Error deleting company", zap.Error(err))
		http.Error(w, "Error deleting company", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ---- Risk Assessment (company-specific) ----

func (hh *HttpHandler) GetRiskAssessment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	assessment, err := hh.store.GetRiskAssessment(ctx, id)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			http.Error(w, "Company not found", http.StatusNotFound)
			return
		}
		zap.L().Error("Error retrieving risk assessment", zap.Error(err))
		http.Error(w, "Error retrieving risk assessment", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(assessment)
}
