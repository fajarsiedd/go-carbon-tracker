package response

import (
	"go-carbon-tracker/entities"
	"go-carbon-tracker/handlers/base"
)

type LoginResponse struct {
	base.Base
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

func (loginResponse LoginResponse) FromEntity(user entities.User) LoginResponse {
	return LoginResponse{
		Base:        loginResponse.Base.FromEntity(user.Base),
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: user.Token,
	}
}
