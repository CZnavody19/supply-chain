package domain

type Location struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Coordinates Coordinates `json:"coordinates"`
	Capacity    int         `json:"capacity"`
}

type InventoryItem struct {
	Product      Product `json:"product"`
	Quantity     int     `json:"quantity"`
	DaysOfSupply int     `json:"daysOfSupply"`
}

type InventoryStatus struct {
	Location Location        `json:"location"`
	Products []InventoryItem `json:"products"`
}
