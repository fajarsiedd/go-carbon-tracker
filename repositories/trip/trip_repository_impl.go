package trip

import (
	"go-carbon-tracker/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type tripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) *tripRepository {
	return &tripRepository{
		db: db,
	}
}

func (repository tripRepository) GetAll() ([]entities.Trip, error) {
	tripModel := ListTrip{}

	if err := repository.db.Find(&tripModel).Error; err != nil {
		return nil, err
	}

	return tripModel.ToListEntity(), nil
}

func (repository tripRepository) GetByID(id string) (entities.Trip, error) {
	tripModel := Trip{}

	if err := repository.db.First(&tripModel, &id).Error; err != nil {
		return entities.Trip{}, err
	}

	return tripModel.ToEntity(), nil
}

func (repository tripRepository) Create(trip entities.Trip) (entities.Trip, error) {
	tripModel := Trip{}.FromEntity(trip)

	err := repository.db.Transaction(func(tx *gorm.DB) error {
		startLocation := tripModel.StartLocation

		if err := tx.Create(&startLocation).Error; err != nil {
			return err
		}

		endLocation := tripModel.StartLocation

		if err := tx.Create(&endLocation).Error; err != nil {
			return err
		}

		tripModel.StartLocationID = startLocation.ID

		tripModel.EndLocationID = endLocation.ID

		tx.Preload(clause.Associations).Create(&tripModel)

		if err := tx.Create(&tripModel).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entities.Trip{}, err
	}

	return tripModel.ToEntity(), nil
}

func (repository tripRepository) Update(trip entities.Trip) (entities.Trip, error) {
	tripModel := Trip{}.FromEntity(trip)

	if err := repository.db.Updates(&tripModel).Error; err != nil {
		return entities.Trip{}, err
	}

	return tripModel.ToEntity(), nil
}

func (repository tripRepository) Delete(id string) error {
	tripModel := Trip{}

	if err := repository.db.Select(clause.Associations).Delete(&tripModel, &id).Error; err != nil {
		return err
	}

	return nil
}
