package vehicle

import (
	"go-carbon-tracker/entities"
	"go-carbon-tracker/models"

	"gorm.io/gorm"
)

type vehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) *vehicleRepository {
	return &vehicleRepository{
		db: db,
	}
}

func (repository vehicleRepository) GetAll(filter entities.Filter) ([]entities.Vehicle, entities.Pagination, error) {
	vehicleModel := models.ListVehicle{}

	query := repository.db.Model(&vehicleModel)

	if filter.Search != "" {
		query = query.Where("name LIKE ?", "%"+filter.Search+"%").Or("vehicle_type LIKE ?", "%"+filter.Search+"%").Or("fuel_type LIKE ?", "%"+filter.Search+"%")
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

	if err := query.Limit(filter.Limit).Offset(offset).Find(&vehicleModel).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	pagination := entities.Pagination{
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalItems: int(totalItems),
		TotalPages: int((int(totalItems) + filter.Limit - 1) / filter.Limit),
	}

	return vehicleModel.ToListEntity(), pagination, nil
}

func (repository vehicleRepository) GetByID(id string) (entities.Vehicle, error) {
	vehicleModel := models.Vehicle{}

	if err := repository.db.First(&vehicleModel, &id).Error; err != nil {
		return entities.Vehicle{}, nil
	}

	return vehicleModel.ToEntity(), nil
}

func (repository vehicleRepository) Create(vehicle entities.Vehicle) (entities.Vehicle, error) {
	vehicleModel := models.Vehicle{}.FromEntity(vehicle)

	if err := repository.db.Create(&vehicleModel).Error; err != nil {
		return entities.Vehicle{}, err
	}

	return vehicleModel.ToEntity(), nil
}

func (repository vehicleRepository) Update(vehicle entities.Vehicle) (entities.Vehicle, error) {
	vehicleModel := models.Vehicle{}.FromEntity(vehicle)

	if err := repository.db.Updates(&vehicleModel).Error; err != nil {
		return entities.Vehicle{}, err
	}

	return vehicleModel.ToEntity(), nil
}

func (repository vehicleRepository) Delete(id string) error {
	vehicleModel := models.Vehicle{}

	if err := repository.db.Delete(&vehicleModel, &id).Error; err != nil {
		return err
	}

	return nil
}
