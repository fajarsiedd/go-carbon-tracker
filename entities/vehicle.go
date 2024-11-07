package entities

import "go-carbon-tracker/constants"

type Vehicle struct {
	Base
	VehicleType    constants.VehicleType
	EmissionFactor float32
	Trips          []Trip
}
