package models

import (
	"go-carbon-tracker/constants/enums"
	"go-carbon-tracker/entities"
)

type Vehicle struct {
	Base
	Name           string            `gorm:"size:191"`
	VehicleType    enums.VehicleType `gorm:"type:enum('MOBIL', 'BUS', 'TRUK', 'MOTOR');column:vehicle_type"`
	FuelType       enums.FuelType    `gorm:"type:enum('SOLAR', 'PREMIUM');column:fuel_type"`
	EmissionFactor float32           `gorm:"type:decimal(10,3)"`
	UserID         string            `gorm:"size:191"`
	Trips          ListTrip
}

type ListVehicle []Vehicle

func (vehicle Vehicle) FromEntity(vehicleEntity entities.Vehicle) Vehicle {
	return Vehicle{
		Base:           vehicle.Base.FromEntity(vehicleEntity.Base),
		Name:           vehicleEntity.Name,
		VehicleType:    vehicleEntity.VehicleType,
		FuelType:       vehicleEntity.FuelType,
		EmissionFactor: vehicleEntity.EmissionFactor,
		Trips:          ListTrip{}.FromListEntity(vehicleEntity.Trips),
		UserID:         vehicleEntity.UserID,
	}
}

func (vehicle Vehicle) ToEntity() entities.Vehicle {
	return entities.Vehicle{
		Base:           vehicle.Base.ToEntity(),
		Name:           vehicle.Name,
		VehicleType:    vehicle.VehicleType,
		FuelType:       vehicle.FuelType,
		EmissionFactor: vehicle.EmissionFactor,
		Trips:          vehicle.Trips.ToListEntity(),
		UserID:         vehicle.UserID,
	}
}

func (ListVehicle) FromListEntity(vehicleEntity []entities.Vehicle) ListVehicle {
	listVehicle := ListVehicle{}

	vehicle := Vehicle{}

	for _, v := range vehicleEntity {
		listVehicle = append(listVehicle, vehicle.FromEntity(v))
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
