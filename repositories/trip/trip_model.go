package trip

import (
	"go-carbon-tracker/entities"
	"go-carbon-tracker/repositories/base"
)

type Location struct {
	base.Base
	Address   string  `json:"address"`
	Latitude  float32 `json:"latitude" gorm:"type:decimal(9,6)"`
	Longitude float32 `json:"longitude" gorm:"type:decimal(9,6)"`
}

type Trip struct {
	base.Base
	DistanceKM      float32  `json:"distance_km" gorm:"type:decimal(10,3)"`
	CarbonEmission  float32  `json:"carbon_emission" gorm:"type:decimal(10,3)"`
	UserID          string   `json:"user_id" gorm:"size:191"`
	VehicleID       string   `json:"vehicle_id" gorm:"size:191"`
	StartLocationID string   `json:"start_location_id" gorm:"size:191"`
	StartLocation   Location `json:"start_location" gorm:"foreignKey:StartLocationID;references:ID"`
	EndLocationID   string   `json:"end_location_id" gorm:"size:191"`
	EndLocation     Location `json:"end_location" gorm:"foreignKey:EndLocationID;references:ID"`
}

type ListTrip []Trip

func (trip Trip) FromEntity(tripEntity entities.Trip) Trip {
	return Trip{
		Base:            trip.Base.FromEntity(tripEntity.Base),
		DistanceKM:      tripEntity.DistanceKM,
		CarbonEmission:  tripEntity.CarbonEmission,
		UserID:          tripEntity.UserID,
		VehicleID:       tripEntity.VehicleID,
		StartLocationID: tripEntity.StartLocationID,
		StartLocation: Location{
			Base:     trip.Base.FromEntity(tripEntity.StartLocation.Base),
			Address:  tripEntity.StartLocation.Address,
			Latitude: tripEntity.StartLocation.Latitude,
		},
		EndLocationID: tripEntity.EndLocationID,
		EndLocation: Location{
			Base:     trip.Base.FromEntity(tripEntity.EndLocation.Base),
			Address:  tripEntity.EndLocation.Address,
			Latitude: tripEntity.EndLocation.Latitude,
		},
	}
}

func (trip Trip) ToEntity() entities.Trip {
	return entities.Trip{
		Base:            trip.Base.ToEntity(),
		DistanceKM:      trip.DistanceKM,
		CarbonEmission:  trip.CarbonEmission,
		UserID:          trip.UserID,
		VehicleID:       trip.VehicleID,
		StartLocationID: trip.StartLocationID,
		StartLocation: entities.Location{
			Base:     trip.StartLocation.Base.ToEntity(),
			Address:  trip.StartLocation.Address,
			Latitude: trip.StartLocation.Latitude,
		},
		EndLocationID: trip.EndLocationID,
		EndLocation: entities.Location{
			Base:     trip.EndLocation.Base.ToEntity(),
			Address:  trip.EndLocation.Address,
			Latitude: trip.EndLocation.Latitude,
		},
	}
}

func (ListTrip) FromListEntity(tripEntity []entities.Trip) ListTrip {
	listTrip := ListTrip{}

	for _, v := range tripEntity {
		trip := Trip{}.FromEntity(v)

		listTrip = append(listTrip, trip)
	}

	return listTrip
}

func (listTrip ListTrip) ToListEntity() []entities.Trip {
	tripEntity := []entities.Trip{}

	for _, v := range listTrip {
		tripEntity = append(tripEntity, v.ToEntity())
	}

	return tripEntity
}
