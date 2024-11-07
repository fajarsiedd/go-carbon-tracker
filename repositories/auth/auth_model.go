package auth

import (
	"go-carbon-tracker/entities"
	"go-carbon-tracker/repositories/base"
	"go-carbon-tracker/repositories/trip"
	"go-carbon-tracker/repositories/vehicle"
)

type User struct {
	base.Base
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Trips    trip.ListTrip
	Vehicles vehicle.ListVehicle
}

func (user User) FromEntity(userEntity entities.User) User {
	return User{
		Base:     user.Base.FromEntity(userEntity.Base),
		Name:     userEntity.Name,
		Email:    userEntity.Email,
		Password: userEntity.Password,
		Trips:    trip.ListTrip{}.FromListEntity(userEntity.Trips),
		Vehicles: vehicle.ListVehicle{}.FromListEntity(userEntity.Vehicles),
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
