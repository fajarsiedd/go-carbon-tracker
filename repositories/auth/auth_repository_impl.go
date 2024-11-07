package auth

import (
	"go-carbon-tracker/entities"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (repository *authRepository) Login(user entities.User) (entities.User, error) {
	userModel := User{}.FromEntity(user)

	result := repository.db.First(&userModel, "email = ?", userModel.Email)

	if result.Error != nil {
		return entities.User{}, result.Error
	}

	return userModel.ToEntity(), nil
}

func (repository authRepository) Register(user entities.User) (entities.User, error) {
	userModel := User{}.FromEntity(user)

	result := repository.db.Create(&userModel)

	if err := result.Error; err != nil {
		return entities.User{}, err
	}

	return userModel.ToEntity(), nil
}
