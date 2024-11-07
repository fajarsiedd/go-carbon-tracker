package request

import (
	"go-carbon-tracker/constants/enums"
	"go-carbon-tracker/entities"
)

type VehicleRequest struct {
	ID          string
	Name        string            `json:"name" validate:"required"`
	VehicleType enums.VehicleType `json:"vehicle_type" validate:"required,oneof=MOBIL MOTOR TRUK BUS"`
	FuelType    enums.FuelType    `json:"fuel_type" validate:"required,oneof=PREMIUM SOLAR"`
}

func (vehicleRequest VehicleRequest) ToEntity() entities.Vehicle {
	return entities.Vehicle{
		Base:        entities.Base{ID: vehicleRequest.ID},
		Name:        vehicleRequest.Name,
		VehicleType: vehicleRequest.VehicleType,
		FuelType:    vehicleRequest.FuelType,
	}
}
