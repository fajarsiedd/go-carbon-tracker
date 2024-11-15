package utils

import (
	c "go-carbon-tracker/constants"
	"net/http"
)

func GetStatusCodeBySuccessMessage(msg string) int {
	switch msg {
	case c.CREATE_TRIP_SUCCESS, c.CREATE_VEHICLE_SUCCESS, c.REGISTER_SUCCESS:
		return http.StatusCreated
	default:
		return http.StatusOK
	}
}
