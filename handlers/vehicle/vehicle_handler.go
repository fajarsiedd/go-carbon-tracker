package vehicle

import (
	"go-carbon-tracker/handlers/base"
	"go-carbon-tracker/handlers/vehicle/request"
	"go-carbon-tracker/handlers/vehicle/response"
	"go-carbon-tracker/middlewares"
	"go-carbon-tracker/usecases/vehicle"

	"github.com/labstack/echo/v4"
)

type vehicleHandler struct {
	usecase vehicle.VehicleUsecase
}

func NewVehicleHandler(u vehicle.VehicleUsecase) *vehicleHandler {
	return &vehicleHandler{
		usecase: u,
	}
}

func (handler vehicleHandler) GetAll(c echo.Context) error {
	result, err := handler.usecase.GetAll()

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, response.ListVehicleResponse{}.FromListEntity(result))
}

func (handler vehicleHandler) GetByID(c echo.Context) error {
	id := c.Param("id")

	result, err := handler.usecase.GetByID(id)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, response.VehicleResponse{}.FromEntity(result))
}

func (handler vehicleHandler) Create(c echo.Context) error {
	req := request.VehicleRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	vehicleReq := req.ToEntity()

	claims, err := middlewares.GetUser(c)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	vehicleReq.UserID = claims.UserID

	result, err := handler.usecase.Create(vehicleReq)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, response.VehicleResponse{}.FromEntity(result))
}

func (handler vehicleHandler) Update(c echo.Context) error {
	id := c.Param("id")

	req := request.VehicleRequest{ID: id}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	vehicleReq := req.ToEntity()

	claims, err := middlewares.GetUser(c)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	vehicleReq.UserID = claims.UserID

	result, err := handler.usecase.Update(vehicleReq)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, response.VehicleResponse{}.FromEntity(result))
}

func (handler vehicleHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := handler.usecase.Delete(id); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, nil)
}
