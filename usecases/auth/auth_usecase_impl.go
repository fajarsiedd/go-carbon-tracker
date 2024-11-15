package auth

import (	
	"go-carbon-tracker/constants"
	"go-carbon-tracker/entities"
	"go-carbon-tracker/middlewares"
	"go-carbon-tracker/repositories/auth"
	"go-carbon-tracker/utils"
)

type authUsecase struct {
	repository auth.AuthRepository
	jwtAuth    *middlewares.JWTConfig
}

func NewAuthUsecase(r auth.AuthRepository, jwtAuth *middlewares.JWTConfig) *authUsecase {
	return &authUsecase{
		repository: r,
		jwtAuth:    jwtAuth,
	}
}

func (usecase authUsecase) Login(user entities.User) (entities.User, error) {
	result, err := usecase.repository.Login(user)

	if err != nil {
		return entities.User{}, err
	}

	match, err := utils.ComparePassword(user.Password, result.Password)

	isFailed := err != nil || !match
	if isFailed {
		return entities.User{}, constants.Err_INVALID_PASSWORD
	}

	token, err := usecase.jwtAuth.GenerateToken(result.ID)
	if err != nil {
		return entities.User{}, err
	}

	result.Token = token

	return result, nil
}

func (usecase authUsecase) Register(user entities.User) (entities.User, error) {
	config := &utils.ArgonConfig{
		Memory:     64 * 1024,
		Iterations: 3,
		Pararelism: 2,
		SaltLength: 16,
		KeyLength:  32,
	}

	var err error
	user.Password, err = utils.CreatePassword(user.Password, config)
	if err != nil {
		return entities.User{}, err
	}

	result, err := usecase.repository.Register(user)

	if err != nil {
		return entities.User{}, err
	}

	return result, nil
}
