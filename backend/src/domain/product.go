package domain

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	SKU      string  `json:"sku"`
	Price    float64 `json:"price"`
	Weight   float64 `json:"weight"`
	LeadTime int     `json:"leadTime"`
	Status   string  `json:"status"`
}

type BOMEntry struct {
	Component Component `json:"component"`
	Quantity  int       `json:"quantity"`
	Position  int       `json:"position"`
}

type BOMDetailedEntry struct {
	Component Component           `json:"component"`
	Quantity  int                 `json:"quantity"`
	Position  int                 `json:"position"`
	Suppliers []ComponentSupplier `json:"suppliers"`
}

type ComponentSupplier struct {
	Company  Company `json:"company"`
	Price    float64 `json:"price"`
	LeadTime int     `json:"leadTime"`
	MinOrder int     `json:"minOrder"`
}

type AlternativeSupplier struct {
	Company     Company `json:"company"`
	Price       float64 `json:"price"`
	Reliability float64 `json:"reliability"`
	LeadTime    int     `json:"leadTime"`
}
