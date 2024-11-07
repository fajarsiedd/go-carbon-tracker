package vehicle

import (
	"go-carbon-tracker/entities"

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

func (repository vehicleRepository) GetAll() ([]entities.Vehicle, error) {
	vehicleModel := ListVehicle{}

	if err := repository.db.Find(&vehicleModel).Error; err != nil {
		return nil, err
	}

	return vehicleModel.ToListEntity(), nil
}

func (repository vehicleRepository) GetByID(id string) (entities.Vehicle, error) {
	vehicleModel := Vehicle{}

	if err := repository.db.First(&vehicleModel, &id).Error; err != nil {
		return entities.Vehicle{}, nil
	}

	return vehicleModel.ToEntity(), nil
}

func (repository vehicleRepository) Create(vehicle entities.Vehicle) (entities.Vehicle, error) {
	vehicleModel := Vehicle{}.FromEntity(vehicle)

	if err := repository.db.Create(&vehicleModel).Error; err != nil {
		return entities.Vehicle{}, err
	}

	return vehicleModel.ToEntity(), nil
}

func (repository vehicleRepository) Update(vehicle entities.Vehicle) (entities.Vehicle, error) {
	vehicleModel := Vehicle{}.FromEntity(vehicle)

	if err := repository.db.Updates(&vehicleModel).Error; err != nil {
		return entities.Vehicle{}, err
	}

	return vehicleModel.ToEntity(), nil
}

func (repository vehicleRepository) Delete(id string) error {
	vehicleModel := Vehicle{}

	if err := repository.db.Delete(&vehicleModel, &id).Error; err != nil {
		return err
	}

	return nil
}
