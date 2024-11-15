package models

import (
	"go-carbon-tracker/entities"
)

type User struct {
	Base
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Trips    ListTrip
	Vehicles ListVehicle
}

func (user User) FromEntity(userEntity entities.User) User {
	return User{
		Base:     user.Base.FromEntity(userEntity.Base),
		Name:     userEntity.Name,
		Email:    userEntity.Email,
		Password: userEntity.Password,
		Trips:    ListTrip{}.FromListEntity(userEntity.Trips),
		Vehicles: ListVehicle{}.FromListEntity(userEntity.Vehicles),
	}
}

func (user User) ToEntity() entities.User {
	return entities.User{
		Base:     user.Base.ToEntity(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Trips:    user.Trips.ToListEntity(),
		Vehicles: user.Vehicles.ToListEntity(),
	}
}
