package response

import (
	"go-carbon-tracker/entities"
	"go-carbon-tracker/handlers/base"
)

type RegisterResponse struct {
	base.Base
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (registerResponse RegisterResponse) FromEntity(user entities.User) RegisterResponse {
	return RegisterResponse{
		Base:  registerResponse.Base.FromEntity(user.Base),
		Name:  user.Name,
		Email: user.Email,
	}
}
