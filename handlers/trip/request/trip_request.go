package request

import "go-carbon-tracker/entities"

type TripRequest struct {
	ID            string
	StartLocation string `json:"start_location" validate:"required"`
	EndLocation   string `json:"end_location" validate:"required"`
	VehicleID     string `json:"vehicle_id" validate:"required"`
}

func (vehicleRequest TripRequest) ToEntity() entities.Trip {
	return entities.Trip{
		Base: entities.Base{ID: vehicleRequest.ID},
		StartLocation: entities.Location{
			Address: vehicleRequest.StartLocation,
		},
		EndLocation: entities.Location{
			Address: vehicleRequest.EndLocation,
		},
		VehicleID: vehicleRequest.VehicleID,
	}
}
