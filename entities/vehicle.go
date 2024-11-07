package entities

import "go-carbon-tracker/constants/enums"

type Vehicle struct {
	Base
	Name           string
	VehicleType    enums.VehicleType
	FuelType       enums.FuelType
	EmissionFactor float32
	UserID         string
	Trips          []Trip
}
