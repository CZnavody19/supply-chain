package domain

type Component struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Criticality string  `json:"criticality"`
}
