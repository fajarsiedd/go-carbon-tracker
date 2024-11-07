package vehicle

import "go-carbon-tracker/entities"

type VehicleRepository interface {
	GetAll() ([]entities.Vehicle, error)
	GetByID(id string) ([]entities.Vehicle, error)
	Create(vehicle entities.Vehicle) (entities.Vehicle, error)
	Update(vehicle entities.Vehicle) (entities.Vehicle, error)
	Delete(id string) error
}
