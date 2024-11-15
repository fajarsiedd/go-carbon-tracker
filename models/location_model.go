package models

type Location struct {
	Base
	Address   string
	Latitude  float64 `gorm:"type:decimal(9,6)"`
	Longitude float64 `gorm:"type:decimal(9,6)"`
}
