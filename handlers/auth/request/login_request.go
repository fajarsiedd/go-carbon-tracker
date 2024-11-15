package request

import "go-carbon-tracker/entities"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (loginRequest LoginRequest) ToEntity() entities.User {
	return entities.User{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}
}
