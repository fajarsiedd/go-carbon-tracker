package response

import (
	"go-carbon-tracker/constants/enums"
	"go-carbon-tracker/entities"
	"go-carbon-tracker/handlers/base"
)

type location struct {
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type vehicle struct {
	base.Base
	Name           string            `json:"name"`
	VehicleType    enums.VehicleType `json:"vehicle_type"`
	FuelType       enums.FuelType    `json:"fuel_type"`
	EmissionFactor float32           `json:"emission_factor"`
}

type user struct {
	base.Base
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TripResponsePopulate struct {
	base.Base
	DistanceKM     int      `json:"distance_km"`
	CarbonEmission float32  `json:"carbon_emission"`
	User           user     `json:"user"`
	Vehicle        vehicle  `json:"vehicle"`
	StartLocation  location `json:"start_location"`
	EndLocation    location `json:"end_location"`
	Tips           string   `json:"tips"`
}

type TripResponse struct {
	base.Base
	DistanceKM     int     `json:"distance_km"`
	CarbonEmission float32 `json:"carbon_emission"`
	UserID         string  `json:"user_id"`
	VehicleID      string  `json:"vehicle_id"`
	Tips           string  `json:"tips"`
}

type ListTripResponse []TripResponse

func (tripResponse TripResponsePopulate) FromEntity(tripEntity entities.Trip) TripResponsePopulate {
	return TripResponsePopulate{
		Base:           tripResponse.Base.FromEntity(tripEntity.Base),
		DistanceKM:     tripEntity.DistanceKM,
		CarbonEmission: tripEntity.CarbonEmission,
		User: user{
			Base:  tripResponse.User.Base.FromEntity(tripEntity.User.Base),
			Name:  tripEntity.User.Name,
			Email: tripEntity.User.Email,
		},
		Vehicle: vehicle{
			Base:           tripResponse.Vehicle.Base.FromEntity(tripEntity.Vehicle.Base),
			Name:           tripEntity.Vehicle.Name,
			VehicleType:    tripEntity.Vehicle.VehicleType,
			FuelType:       tripEntity.Vehicle.FuelType,
			EmissionFactor: tripEntity.Vehicle.EmissionFactor,
		},
		StartLocation: location{
			Address:   tripEntity.StartLocation.Address,
			Latitude:  tripEntity.StartLocation.Latitude,
			Longitude: tripEntity.StartLocation.Longitude,
		},
		EndLocation: location{
			Address:   tripEntity.EndLocation.Address,
			Latitude:  tripEntity.EndLocation.Latitude,
			Longitude: tripEntity.EndLocation.Longitude,
		},
		Tips: tripEntity.Tips,
	}
}

func (ListTripResponse) FromListEntity(tripEntity []entities.Trip) ListTripResponse {
	listTrip := ListTripResponse{}

	trip := TripResponse{}

	for _, v := range tripEntity {
		trip.Base = trip.Base.FromEntity(v.Base)
		trip.DistanceKM = v.DistanceKM
		trip.CarbonEmission = v.CarbonEmission
		trip.UserID = v.UserID
		trip.VehicleID = v.VehicleID
		trip.Tips = v.Tips

		listTrip = append(listTrip, trip)
	}

	return listTrip
}
