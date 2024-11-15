package vehicle

import (
	"go-carbon-tracker/entities"
	"go-carbon-tracker/repositories/vehicle"
	"go-carbon-tracker/utils"
)

type vehicleUsecase struct {
	repository vehicle.VehicleRepository
}

func NewVehicleUsecase(r vehicle.VehicleRepository) *vehicleUsecase {
	return &vehicleUsecase{
		repository: r,
	}
}

func (usecase vehicleUsecase) GetAll(filter entities.Filter) ([]entities.Vehicle, entities.Pagination, error) {
	return usecase.repository.GetAll(filter)
}

func (usecase vehicleUsecase) GetByID(id string) (entities.Vehicle, error) {
	return usecase.repository.GetByID(id)
}

func (usecase vehicleUsecase) Create(vehicle entities.Vehicle) (entities.Vehicle, error) {
	co2Factor := utils.GetCO2Factor(vehicle.FuelType)

	energyConsume := utils.GetEnergyConsume(vehicle.VehicleType)

	vehicle.EmissionFactor = co2Factor * energyConsume

	return usecase.repository.Create(vehicle)
}

func (usecase vehicleUsecase) Update(vehicle entities.Vehicle) (entities.Vehicle, error) {
	co2Factor := utils.GetCO2Factor(vehicle.FuelType)

	energyConsume := utils.GetEnergyConsume(vehicle.VehicleType)

	vehicle.EmissionFactor = co2Factor * energyConsume

	return usecase.repository.Update(vehicle)
}

func (usecase vehicleUsecase) Delete(id string) error {
	return usecase.repository.Delete(id)
}
