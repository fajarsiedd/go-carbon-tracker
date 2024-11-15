package trip

import (
	"go-carbon-tracker/entities"
	"go-carbon-tracker/models"

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

func (repository tripRepository) GetAll(filter entities.Filter) ([]entities.Trip, entities.Pagination, error) {
	tripModel := models.ListTrip{}

	query := repository.db.Model(&tripModel)

	if filter.Search != "" {
		query = query.Where("tips LIKE ?", "%"+filter.Search+"%")
	}

	if !filter.StartDate.IsZero() && !filter.EndDate.IsZero() {
		query = query.Where("created_at BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
	}

	query = query.Order(filter.SortBy + " " + filter.Sort)

	query.Where("user_id = ?", filter.UserID)

	var totalItems int64

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	offset := (filter.Page - 1) * filter.Limit

	if err := query.Limit(filter.Limit).Offset(offset).Find(&tripModel).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	pagination := entities.Pagination{
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalItems: int(totalItems),
		TotalPages: int((int(totalItems) + filter.Limit - 1) / filter.Limit),
	}

	return tripModel.ToListEntity(), pagination, nil
}

func (repository tripRepository) GetByID(id string) (entities.Trip, error) {
	tripModel := models.Trip{}

	if err := repository.db.Preload(clause.Associations).First(&tripModel, &id).Error; err != nil {
		return entities.Trip{}, err
	}

	return tripModel.ToEntity(), nil
}

func (repository tripRepository) Create(trip entities.Trip) (entities.Trip, error) {
	tripModel := models.Trip{}.FromEntity(trip)

	if err := repository.db.Omit("Vehicle", "User").Create(&tripModel).Preload(clause.Associations).Find(&tripModel).Error; err != nil {
		return entities.Trip{}, err
	}

	return tripModel.ToEntity(), nil
}

func (repository tripRepository) Update(trip entities.Trip) (entities.Trip, error) {
	tripModel := models.Trip{}.FromEntity(trip)

	err := repository.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Vehicle", "User", "StartLocation", "EndLocation").
			Updates(&tripModel).Error; err != nil {
			return err
		}

		if err := tx.Model(&tripModel.StartLocation).Where("id = ?", tripModel.StartLocationID).
			Updates(tripModel.StartLocation).Error; err != nil {
			return err
		}

		if err := tx.Model(&tripModel.EndLocation).Where("id = ?", tripModel.EndLocationID).
			Updates(tripModel.EndLocation).Error; err != nil {
			return err
		}

		if err := tx.Preload(clause.Associations).Find(&tripModel).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entities.Trip{}, err
	}

	return tripModel.ToEntity(), nil
}

func (repository tripRepository) Delete(id string) error {
	tripModel := models.Trip{Base: models.Base{ID: id}}

	if err := repository.db.Select("StartLocation", "EndLocation").Delete(&tripModel).Error; err != nil {
		return err
	}

	return nil
}
