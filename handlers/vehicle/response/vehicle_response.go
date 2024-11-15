package response

import (
	"go-carbon-tracker/constants/enums"
	"go-carbon-tracker/entities"
	"go-carbon-tracker/handlers/base"
)

type VehicleResponse struct {
	base.Base
	Name           string            `json:"name"`
	VehicleType    enums.VehicleType `json:"vehicle_type"`
	FuelType       enums.FuelType    `json:"fuel_type"`
	EmissionFactor float32           `json:"emission_factor"`
	UserID         string            `json:"user_id"`
}

type ListVehicleResponse []VehicleResponse

func (vehicleResponse VehicleResponse) FromEntity(vehicleEntity entities.Vehicle) VehicleResponse {
	return VehicleResponse{
		Base:           vehicleResponse.Base.FromEntity(vehicleEntity.Base),
		Name:           vehicleEntity.Name,
		VehicleType:    vehicleEntity.VehicleType,
		FuelType:       vehicleEntity.FuelType,
		EmissionFactor: vehicleEntity.EmissionFactor,
		UserID:         vehicleEntity.UserID,
	}
}

func (ListVehicleResponse) FromListEntity(vehicleEntity []entities.Vehicle) ListVehicleResponse {
	listVehicle := ListVehicleResponse{}

	vehicle := VehicleResponse{}

	for _, v := range vehicleEntity {
		listVehicle = append(listVehicle, vehicle.FromEntity(v))
	}

	return listVehicle
}
