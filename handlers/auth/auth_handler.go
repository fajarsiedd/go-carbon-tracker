package auth

import (
	"go-carbon-tracker/constants"
	"go-carbon-tracker/handlers/auth/request"
	"go-carbon-tracker/handlers/auth/response"
	"go-carbon-tracker/handlers/base"
	"go-carbon-tracker/usecases/auth"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	usecase auth.AuthUsecase
}

func NewAuthHandler(usecase auth.AuthUsecase) *authHandler {
	return &authHandler{
		usecase: usecase,
	}
}

func (handler authHandler) Login(c echo.Context) error {
	req := request.LoginRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	result, err := handler.usecase.Login(req.ToEntity())

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, constants.LOGIN_SUCCESS, response.LoginResponse{}.FromEntity(result))
}

func (handler authHandler) Register(c echo.Context) error {
	req := request.RegisterRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	result, err := handler.usecase.Register(req.ToEntity())

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, constants.REGISTER_SUCCESS, response.RegisterResponse{}.FromEntity(result))
}
