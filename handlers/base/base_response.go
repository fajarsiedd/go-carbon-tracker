package base

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func SuccesResponse(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "success",
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, BaseResponse{
		Status:  false,
		Message: err.Error(),
	})
}
