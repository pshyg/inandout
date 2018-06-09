package models

// Location has user's latitude and longitude
type Location struct {
	UserID    string `json:"userid"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
