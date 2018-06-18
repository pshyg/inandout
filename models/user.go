package models

// User save account data
type User struct {
	ID        string `json:"userID" query:"userID" param:"userID" gorm:"primary_key"`
	Password  string `json:"password"`
	Locations []Location
}
