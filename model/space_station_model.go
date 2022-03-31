package model

// URL: http://api.open-notify.org/iss-now.json
// Example data:
// {"iss_position": {"longitude": "31.0058", "latitude": "-2.5085"}, "message": "success", "timestamp": 1648548448}
type SpaceStationData struct {
	PositionData IssPosition `json:"iss_position"`
}

type IssPosition struct {
	Longitude float64 `json:"longitude,string"`
	Latitude  float64 `json:"latitude,string"`
}
