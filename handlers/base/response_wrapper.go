package base

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseWrapper struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func SuccesResponse(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, ResponseWrapper{
		Status:  true,
		Message: "success",
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, ResponseWrapper{
		Status:  false,
		Message: err.Error(),
	})
}
