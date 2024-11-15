package models

import (
	"go-carbon-tracker/entities"
)

type Trip struct {
	Base
	DistanceKM      int
	CarbonEmission  float32  `gorm:"type:decimal(10,3)"`
	UserID          string   `gorm:"size:191"`
	User            User     `gorm:"foreignKey:UserID;references:ID"`
	VehicleID       string   `gorm:"size:191"`
	Vehicle         Vehicle  `gorm:"foreignKey:VehicleID;references:ID"`
	StartLocationID string   `gorm:"size:191"`
	StartLocation   Location `gorm:"foreignKey:StartLocationID;references:ID"`
	EndLocationID   string   `gorm:"size:191"`
	EndLocation     Location `gorm:"foreignKey:EndLocationID;references:ID"`
	Tips            string
}

type ListTrip []Trip

func (trip Trip) FromEntity(tripEntity entities.Trip) Trip {
	return Trip{
		Base:            trip.Base.FromEntity(tripEntity.Base),
		DistanceKM:      tripEntity.DistanceKM,
		CarbonEmission:  tripEntity.CarbonEmission,
		UserID:          tripEntity.UserID,
		User:            trip.User.FromEntity(tripEntity.User),
		VehicleID:       tripEntity.VehicleID,
		Vehicle:         trip.Vehicle.FromEntity(tripEntity.Vehicle),
		StartLocationID: tripEntity.StartLocationID,
		StartLocation: Location{
			Base:      trip.Base.FromEntity(tripEntity.StartLocation.Base),
			Address:   tripEntity.StartLocation.Address,
			Latitude:  tripEntity.StartLocation.Latitude,
			Longitude: tripEntity.StartLocation.Longitude,
		},
		EndLocationID: tripEntity.EndLocationID,
		EndLocation: Location{
			Base:      trip.Base.FromEntity(tripEntity.EndLocation.Base),
			Address:   tripEntity.EndLocation.Address,
			Latitude:  tripEntity.EndLocation.Latitude,
			Longitude: tripEntity.EndLocation.Longitude,
		},
		Tips: tripEntity.Tips,
	}
}

func (trip Trip) ToEntity() entities.Trip {
	return entities.Trip{
		Base:            trip.Base.ToEntity(),
		DistanceKM:      trip.DistanceKM,
		CarbonEmission:  trip.CarbonEmission,
		UserID:          trip.UserID,
		User:            trip.User.ToEntity(),
		VehicleID:       trip.VehicleID,
		Vehicle:         trip.Vehicle.ToEntity(),
		StartLocationID: trip.StartLocationID,
		StartLocation: entities.Location{
			Base:      trip.StartLocation.Base.ToEntity(),
			Address:   trip.StartLocation.Address,
			Latitude:  trip.StartLocation.Latitude,
			Longitude: trip.StartLocation.Longitude,
		},
		EndLocationID: trip.EndLocationID,
		EndLocation: entities.Location{
			Base:      trip.EndLocation.Base.ToEntity(),
			Address:   trip.EndLocation.Address,
			Latitude:  trip.EndLocation.Latitude,
			Longitude: trip.EndLocation.Longitude,
		},
		Tips: trip.Tips,
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
