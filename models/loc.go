package models

import "fmt"

// Locations has multi-value location data for GetAllOpponentLoc method
type Locations struct {
	Locs []*Location `json:"locs"`
}

// Location has user's latitude and longitude
type Location struct {
	UserID    string `json:"userID" query:"userID" param:"userID" gorm:"primary_key"`
	Time      int    `json:"time" gorm:"primary_key" sql:"auto_increment:false"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (l *Location) String() string {
	return fmt.Sprintf("%d\r\n%s\r\n%s", l.Time, l.Latitude, l.Longitude)
}
