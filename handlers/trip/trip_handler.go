package trip

import (
	"go-carbon-tracker/constants"
	"go-carbon-tracker/handlers/base"
	"go-carbon-tracker/handlers/trip/request"
	"go-carbon-tracker/handlers/trip/response"
	"go-carbon-tracker/middlewares"
	"go-carbon-tracker/usecases/trip"
	"go-carbon-tracker/utils"

	"github.com/labstack/echo/v4"
)

type tripHandler struct {
	usecase trip.TripUsecase
}

func NewTripHandler(u trip.TripUsecase) *tripHandler {
	return &tripHandler{
		usecase: u,
	}
}

func (handler tripHandler) GetAll(c echo.Context) error {
	filter, err := utils.GetFilter(c)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	claims, err := middlewares.GetUser(c)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	filter.UserID = claims.UserID

	result, pagination, err := handler.usecase.GetAll(filter)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponsePagination(c, constants.GET_ALL_TRIPS_SUCCESS, pagination, response.ListTripResponse{}.FromListEntity(result))
}

func (handler tripHandler) GetByID(c echo.Context) error {
	id := c.Param("id")

	result, err := handler.usecase.GetByID(id)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, constants.GET_TRIP_SUCCESS, response.TripResponsePopulate{}.FromEntity(result))
}

func (handler tripHandler) Create(c echo.Context) error {
	req := request.TripRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	tripReq := req.ToEntity()

	claims, err := middlewares.GetUser(c)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	tripReq.UserID = claims.UserID

	result, err := handler.usecase.Create(tripReq)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, constants.CREATE_TRIP_SUCCESS, response.TripResponsePopulate{}.FromEntity(result))
}

func (handler tripHandler) Update(c echo.Context) error {
	id := c.Param("id")

	req := request.TripRequest{ID: id}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	tripReq := req.ToEntity()

	claims, err := middlewares.GetUser(c)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	tripReq.UserID = claims.UserID

	result, err := handler.usecase.Update(tripReq)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, constants.UPDATE_TRIP_SUCCESS, response.TripResponsePopulate{}.FromEntity(result))
}

func (handler tripHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := handler.usecase.Delete(id); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, constants.DELETE_TRIP_SUCCESS, nil)
}
