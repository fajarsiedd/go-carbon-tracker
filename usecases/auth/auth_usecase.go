package auth

import "go-carbon-tracker/entities"

type AuthUsecase interface {
	Login(user entities.User) (entities.User, error)
	Register(user entities.User) (entities.User, error)
}
