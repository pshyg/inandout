package models

import "fmt"

// Location has user's latitude and longitude
type Location struct {
	UserID    string `json:"userid"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Time      int    `json:"time"`
}

func (l *Location) String() string {
	return fmt.Sprintf("%d\r\n%s\r\n%s", l.Time, l.Latitude, l.Longitude)
}
