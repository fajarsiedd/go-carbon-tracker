package trip

import "go-carbon-tracker/entities"

type TripRepository interface {
	GetAll() ([]entities.Trip, error)
	GetByID(id string) ([]entities.Trip, error)
	Create(trip entities.Trip) (entities.Trip, error)
	Update(trip entities.Trip) (entities.Trip, error)
	Delete(id string) error
}
