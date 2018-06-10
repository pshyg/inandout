package models

import "fmt"

// Location has user's latitude and longitude
type Location struct {
	UserID    string `json:"userid" gorm:"primary_key"`
	Time      int    `json:"time" gorm:"primary_key" sql:"auto_increment:false"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (l *Location) String() string {
	return fmt.Sprintf("%d\r\n%s\r\n%s", l.Time, l.Latitude, l.Longitude)
}
