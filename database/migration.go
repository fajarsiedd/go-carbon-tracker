package database

import (
	"go-carbon-tracker/repositories/auth"
	"go-carbon-tracker/repositories/trip"
	"go-carbon-tracker/repositories/vehicle"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&auth.User{}, &trip.Location{}, &trip.Trip{}, &vehicle.Vehicle{})

	// Add table suffix when creating tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&auth.User{}, &trip.Location{}, &trip.Trip{}, &vehicle.Vehicle{})
}
