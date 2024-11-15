package entities

type Trip struct {
	Base
	DistanceKM      int
	CarbonEmission  float32
	UserID          string
	User            User
	VehicleID       string
	Vehicle         Vehicle
	StartLocationID string
	StartLocation   Location
	EndLocationID   string
	EndLocation     Location
	Tips            string
}
