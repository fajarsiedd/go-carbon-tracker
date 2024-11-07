package auth

import (
	"go-carbon-tracker/entities"
	"go-carbon-tracker/repositories/base"
	"go-carbon-tracker/repositories/trip"
)

type User struct {
	base.Base
	Name     string        `json:"name"`
	Email    string        `json:"email" gorm:"unique"`
	Password string        `json:"password"`
	Trips    trip.ListTrip `json:"trips,omitempty"`
}

func (user User) FromEntity(userEntity entities.User) User {
	return User{
		Base:     user.Base.FromEntity(userEntity.Base),
		Name:     userEntity.Name,
		Email:    userEntity.Email,
		Password: userEntity.Password,
	}
}

func (user User) ToEntity() entities.User {
	return entities.User{
		Base:     user.Base.ToEntity(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
