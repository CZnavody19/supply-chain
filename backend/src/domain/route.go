package domain

type Route struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Distance      float64 `json:"distance"`
	EstimatedTime float64 `json:"estimatedTime"`
	Cost          float64 `json:"cost"`
	Reliability   float64 `json:"reliability"`
}

type OptimalRouteResult struct {
	Segments         []RouteSegment `json:"segments"`
	TotalDistance    float64        `json:"totalDistance"`
	TotalTime        float64        `json:"totalTime"`
	TotalCost        float64        `json:"totalCost"`
	TotalReliability float64        `json:"totalReliability"`
}

type RouteSegment struct {
	From     string  `json:"from"`
	To       string  `json:"to"`
	Distance float64 `json:"distance"`
	Time     float64 `json:"time"`
	Cost     float64 `json:"cost"`
}
