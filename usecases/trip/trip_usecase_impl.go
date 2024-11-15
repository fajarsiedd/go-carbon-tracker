package trip

import (
	"go-carbon-tracker/entities"
	"go-carbon-tracker/repositories/trip"
	"go-carbon-tracker/repositories/vehicle"
	"go-carbon-tracker/utils"
)

type tripUsecase struct {
	tripRepository    trip.TripRepository
	vehicleRepository vehicle.VehicleRepository
}

func NewTripUsecase(tripRepository trip.TripRepository, vehicleRepository vehicle.VehicleRepository) *tripUsecase {
	return &tripUsecase{
		tripRepository:    tripRepository,
		vehicleRepository: vehicleRepository,
	}
}

func (usecase tripUsecase) GetAll(filter entities.Filter) ([]entities.Trip, entities.Pagination, error) {
	return usecase.tripRepository.GetAll(filter)
}

func (usecase tripUsecase) GetByID(id string) (entities.Trip, error) {
	return usecase.tripRepository.GetByID(id)
}

func (usecase tripUsecase) Create(trip entities.Trip) (entities.Trip, error) {
	vehicle, err := usecase.vehicleRepository.GetByID(trip.VehicleID)

	if err != nil {
		return entities.Trip{}, err
	}

	// Get Coordinates
	lat1, lng1, err := utils.GetCoordinates(trip.StartLocation.Address)

	if err != nil {
		return entities.Trip{}, err
	}

	trip.StartLocation.Latitude = lat1
	trip.StartLocation.Longitude = lng1

	lat2, lng2, err := utils.GetCoordinates(trip.EndLocation.Address)

	if err != nil {
		return entities.Trip{}, err
	}

	trip.EndLocation.Latitude = lat2
	trip.EndLocation.Longitude = lng2

	// Get Distances
	distance, err := utils.GetDistances(lat1, lng1, lat2, lng2)

	if err != nil {
		return entities.Trip{}, err
	}

	trip.DistanceKM = distance

	trip.CarbonEmission = vehicle.EmissionFactor * float32(distance)

	// Get Tips From Gemini
	trip.Tips = utils.GetTipsFromGemini(trip)

	return usecase.tripRepository.Create(trip)
}

func (usecase tripUsecase) Update(trip entities.Trip) (entities.Trip, error) {
	tripData, err := usecase.tripRepository.GetByID(trip.ID)

	if err != nil {
		return entities.Trip{}, err
	}

	// Get Coordinates
	lat1, lng1, err := utils.GetCoordinates(trip.StartLocation.Address)

	if err != nil {
		return entities.Trip{}, err
	}

	tripData.StartLocation.Address = trip.StartLocation.Address
	tripData.StartLocation.Latitude = lat1
	tripData.StartLocation.Longitude = lng1

	lat2, lng2, err := utils.GetCoordinates(trip.EndLocation.Address)

	if err != nil {
		return entities.Trip{}, err
	}

	tripData.EndLocation.Address = trip.EndLocation.Address
	tripData.EndLocation.Latitude = lat2
	tripData.EndLocation.Longitude = lng2

	// Get Distances
	distance, err := utils.GetDistances(lat1, lng1, lat2, lng2)

	if err != nil {
		return entities.Trip{}, err
	}

	tripData.DistanceKM = distance

	tripData.CarbonEmission = tripData.Vehicle.EmissionFactor * float32(distance)

	// Get Tips From Gemini
	tripData.Tips = utils.GetTipsFromGemini(tripData)

	return usecase.tripRepository.Update(tripData)
}

func (usecase tripUsecase) Delete(id string) error {
	return usecase.tripRepository.Delete(id)
}
