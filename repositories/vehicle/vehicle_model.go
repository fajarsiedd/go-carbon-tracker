package vehicle

import (
	"go-carbon-tracker/constants"
	"go-carbon-tracker/entities"
	"go-carbon-tracker/repositories/base"
	"go-carbon-tracker/repositories/trip"
)

type Vehicle struct {
	base.Base
	VehicleType    constants.VehicleType `json:"vehicle_type" gorm:"type:enum('MOBIL', 'BUS', 'TRUK', 'MOTOR');column:vehicle_type"`
	EmissionFactor float32               `json:"emission_factor" gorm:"type:decimal(10,3)"`
	Trips          trip.ListTrip         `json:"trips,omitempty"`
}

type ListVehicle []Vehicle

func (vehicle Vehicle) FromEntity(vehicleEntity entities.Vehicle) Vehicle {
	return Vehicle{
		Base:           vehicle.Base.FromEntity(vehicleEntity.Base),
		VehicleType:    constants.VehicleType(vehicleEntity.VehicleType),
		EmissionFactor: vehicleEntity.EmissionFactor,
		Trips:          trip.ListTrip{}.FromListEntity(vehicleEntity.Trips),
	}
}

func (vehicle Vehicle) ToEntity() entities.Vehicle {
	return entities.Vehicle{
		Base:           vehicle.Base.ToEntity(),
		VehicleType:    constants.VehicleType(vehicle.VehicleType),
		EmissionFactor: vehicle.EmissionFactor,
		Trips:          vehicle.Trips.ToListEntity(),
	}
}

func FromListEntity(vehicleEntity []entities.Vehicle) ListVehicle {
	listVehicle := ListVehicle{}

	for _, v := range vehicleEntity {
		vehicle := Vehicle{}.FromEntity(v)

		listVehicle = append(listVehicle, vehicle)
	}

	return listVehicle
}

func (listVehicle ListVehicle) ToListEntity() []entities.Vehicle {
	vehicleEntity := []entities.Vehicle{}

	for _, v := range listVehicle {
		vehicleEntity = append(vehicleEntity, v.ToEntity())
	}

	return vehicleEntity
}
