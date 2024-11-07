package entities

type Trip struct {
	Base
	DistanceKM      float32
	CarbonEmission  float32
	UserID          string
	VehicleID       string
	StartLocationID string
	StartLocation   Location
	EndLocationID   string
	EndLocation     Location
}
