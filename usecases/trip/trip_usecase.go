package trip

import "go-carbon-tracker/entities"

type TripUsecase interface {
	GetAll(filter entities.Filter) ([]entities.Trip, entities.Pagination, error)
	GetByID(id string) (entities.Trip, error)
	Create(vehicle entities.Trip) (entities.Trip, error)
	Update(vehicle entities.Trip) (entities.Trip, error)
	Delete(id string) error
}
