package setup

import (
	"github.com/CZnavody19/supply-chain/src/http"
	"github.com/gorilla/mux"
)

func SetupHTTPHandlers(router *mux.Router, handler *http.HttpHandler) {
	router.HandleFunc("/", handler.Ping).Methods("GET")

	// ---- Products CRUD ----
	router.HandleFunc("/api/products", handler.GetProducts).Methods("GET")
	router.HandleFunc("/api/products", handler.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products", handler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products", handler.DeleteProduct).Methods("DELETE")

	// ---- Products BOM ----
	router.HandleFunc("/api/products/bom", handler.GetProductBOM).Methods("GET")
	router.HandleFunc("/api/products/bom/detailed", handler.GetProductBOMDetailed).Methods("GET")
	router.HandleFunc("/api/products/bom", handler.AddBOMComponent).Methods("POST")
	router.HandleFunc("/api/products/bom", handler.UpdateBOMComponent).Methods("PUT")

	// ---- Products Alternative Suppliers ----
	router.HandleFunc("/api/products/alternative-suppliers", handler.GetAlternativeSuppliers).Methods("GET")

	// ---- Companies CRUD ----
	router.HandleFunc("/api/companies", handler.GetCompanies).Methods("GET")
	router.HandleFunc("/api/companies", handler.CreateCompany).Methods("POST")
	router.HandleFunc("/api/companies", handler.UpdateCompany).Methods("PUT")
	router.HandleFunc("/api/companies", handler.DeleteCompany).Methods("DELETE")

	// ---- Companies Risk Assessment ----
	router.HandleFunc("/api/companies/risk-assessment", handler.GetRiskAssessment).Methods("GET")

	// ---- Components CRUD ----
	router.HandleFunc("/api/components", handler.GetComponents).Methods("GET")
	router.HandleFunc("/api/components", handler.CreateComponent).Methods("POST")
	router.HandleFunc("/api/components", handler.UpdateComponent).Methods("PUT")
	router.HandleFunc("/api/components", handler.DeleteComponent).Methods("DELETE")

	// ---- Orders CRUD ----
	router.HandleFunc("/api/orders", handler.GetOrders).Methods("GET")
	router.HandleFunc("/api/orders", handler.CreateOrder).Methods("POST")
	router.HandleFunc("/api/orders/status", handler.UpdateOrderStatus).Methods("PUT")

	// ---- Orders Supply Path & Cost ----
	router.HandleFunc("/api/orders/supply-path", handler.GetSupplyPath).Methods("GET")
	router.HandleFunc("/api/orders/cost-breakdown", handler.GetCostBreakdown).Methods("GET")

	// ---- Locations ----
	router.HandleFunc("/api/locations", handler.GetLocations).Methods("GET")
	router.HandleFunc("/api/locations", handler.CreateLocation).Methods("POST")
	router.HandleFunc("/api/locations/inventory-status", handler.GetInventoryStatus).Methods("GET")

	// ---- Routes ----
	router.HandleFunc("/api/routes/optimal", handler.GetOptimalRoute).Methods("GET")

	// ---- Analytics ----
	router.HandleFunc("/api/analytics/supply-chain-health", handler.GetSupplyChainHealth).Methods("GET")
	router.HandleFunc("/api/analytics/impact-analysis", handler.GetImpactAnalysis).Methods("GET")
	router.HandleFunc("/api/analytics/forecast-delays", handler.GetForecastDelays).Methods("GET")
	router.HandleFunc("/api/analytics/stock-levels", handler.GetStockLevelForecast).Methods("GET")
}
