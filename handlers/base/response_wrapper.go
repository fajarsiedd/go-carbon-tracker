package base

import (
	"go-carbon-tracker/entities"
	"go-carbon-tracker/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Meta struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

type ResponseWrapper struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data,omitempty"`
}

type ResponsePaginationWrapper struct {
	Meta       Meta       `json:"meta"`
	Data       any        `json:"data,omitempty"`
	Pagination Pagination `json:"pagination"`
}

func SuccesResponse(c echo.Context, message string, data any) error {
	statusCode := utils.GetStatusCodeBySuccessMessage(message)

	return c.JSON(statusCode, ResponseWrapper{
		Meta: Meta{
			Status:  true,
			Code:    statusCode,
			Message: message,
		},
		Data: data,
	})
}

func SuccesResponsePagination(c echo.Context, message string, pagination entities.Pagination, data any) error {
	statusCode := utils.GetStatusCodeBySuccessMessage(message)

	return c.JSON(statusCode, ResponsePaginationWrapper{
		Meta: Meta{
			Status:  true,
			Code:    statusCode,
			Message: message,
		},
		Data: data,
		Pagination: Pagination{
			Page:       pagination.Page,
			Limit:      pagination.Limit,
			TotalPages: pagination.TotalPages,
			TotalItems: pagination.TotalItems,
		},
	})
}

func ErrorResponse(c echo.Context, err error) error {
	statusCode := http.StatusInternalServerError

	if errMsg := err.Error(); strings.Contains(errMsg, "validation") {
		statusCode = http.StatusBadRequest
	}

	return c.JSON(statusCode, ResponseWrapper{
		Meta: Meta{
			Status: false,
			Code:   statusCode,
			Error:  err.Error(),
		},
	})
}
