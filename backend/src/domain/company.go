package domain

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Company struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Country     string      `json:"country"`
	Coordinates Coordinates `json:"coordinates"`
	Reliability float64     `json:"reliability"`
}
