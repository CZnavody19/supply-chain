package domain

type Order struct {
	ID        string  `json:"id"`
	OrderDate string  `json:"orderDate"`
	DueDate   string  `json:"dueDate"`
	Quantity  int     `json:"quantity"`
	Status    string  `json:"status"`
	Cost      float64 `json:"cost"`
}
