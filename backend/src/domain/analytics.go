package domain

// ---- Supply Path ----

type SupplyPathResponse struct {
	OrderID       string            `json:"orderId"`
	Product       string            `json:"product"`
	Quantity      int               `json:"quantity"`
	TotalCost     float64           `json:"totalCost"`
	Path          []SupplyPathStage `json:"path"`
	TotalDuration string            `json:"totalDuration"`
	RiskFactors   []string          `json:"riskFactors"`
}

type SupplyPathStage struct {
	Stage    int              `json:"stage"`
	Name     string           `json:"name"`
	Company  *CompanySummary  `json:"company,omitempty"`
	Location *LocationSummary `json:"location,omitempty"`
	From     string           `json:"from,omitempty"`
	To       string           `json:"to,omitempty"`
	Route    *RouteSummary    `json:"route,omitempty"`
	DueDate  string           `json:"dueDate"`
	Status   string           `json:"status"`
}

type CompanySummary struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Reliability float64 `json:"reliability"`
}

type LocationSummary struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country,omitempty"`
}

type RouteSummary struct {
	Distance float64 `json:"distance"`
	Time     string  `json:"time"`
	Cost     float64 `json:"cost"`
}

// ---- Risk Assessment ----

type RiskAssessment struct {
	SupplierID      string            `json:"supplierId"`
	Company         string            `json:"company"`
	RiskScore       float64           `json:"riskScore"`
	Factors         RiskFactors       `json:"factors"`
	CriticalFor     []CriticalProduct `json:"criticalFor"`
	Recommendations []string          `json:"recommendations"`
}

type RiskFactors struct {
	ReliabilityScore   float64 `json:"reliabilityScore"`
	OnTimeDeliveryRate float64 `json:"onTimeDeliveryRate"`
	QualityIssues      float64 `json:"qualityIssues"`
	GeopoliticalRisk   float64 `json:"geopoliticalRisk"`
	FinancialStability float64 `json:"financialStability"`
}

type CriticalProduct struct {
	Product      string `json:"product"`
	Impact       string `json:"impact"`
	Alternatives int    `json:"alternatives"`
}

// ---- Supply Chain Health ----

type SupplyChainHealth struct {
	CriticalComponents []CriticalComponentInfo `json:"criticalComponents"`
	Bottlenecks        []BottleneckInfo        `json:"bottlenecks"`
	HighRiskSuppliers  []HighRiskSupplierInfo  `json:"highRiskSuppliers"`
	Recommendations    []string                `json:"recommendations"`
}

type CriticalComponentInfo struct {
	ComponentID   string `json:"componentId"`
	ComponentName string `json:"componentName"`
	Criticality   string `json:"criticality"`
	SupplierCount int    `json:"supplierCount"`
}

type BottleneckInfo struct {
	LocationID   string  `json:"locationId"`
	LocationName string  `json:"locationName"`
	Utilization  float64 `json:"utilization"`
}

type HighRiskSupplierInfo struct {
	CompanyID   string  `json:"companyId"`
	CompanyName string  `json:"companyName"`
	Reliability float64 `json:"reliability"`
}

// ---- Impact Analysis ----

type ImpactAnalysis struct {
	SupplierID   string       `json:"supplierId"`
	SupplierName string       `json:"supplierName"`
	Impact       ImpactDetail `json:"impact"`
}

type ImpactDetail struct {
	AffectedProducts []AffectedProduct `json:"affectedProducts"`
	EstimatedCost    float64           `json:"estimatedCost"`
	AffectedRevenue  float64           `json:"affectedRevenue"`
	Mitigation       []string          `json:"mitigation"`
}

type AffectedProduct struct {
	ProductID      string `json:"productId"`
	ProductName    string `json:"productName"`
	AffectedOrders int    `json:"affectedOrders"`
	DelayDays      int    `json:"delayDays"`
}

// ---- Cost Breakdown ----

type CostBreakdown struct {
	OrderID           string  `json:"orderId"`
	MaterialCost      float64 `json:"materialCost"`
	ManufacturingCost float64 `json:"manufacturingCost"`
	LogisticsCost     float64 `json:"logisticsCost"`
	TotalCost         float64 `json:"totalCost"`
}

// ---- Forecast / Predictions ----

type ForecastDelay struct {
	ProductID   string  `json:"productId"`
	ProductName string  `json:"productName"`
	AvgDelay    float64 `json:"avgDelayDays"`
	Probability float64 `json:"probability"`
	RiskLevel   string  `json:"riskLevel"`
}

type StockLevelForecast struct {
	ProductID    string            `json:"productId"`
	ProductName  string            `json:"productName"`
	CurrentStock int               `json:"currentStock"`
	Projections  []StockProjection `json:"projections"`
}

type StockProjection struct {
	Month          string `json:"month"`
	ProjectedStock int    `json:"projectedStock"`
	Status         string `json:"status"`
}
