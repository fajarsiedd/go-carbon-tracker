package database

import (
	"go-carbon-tracker/models"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Location{}, &models.Trip{}, &models.Vehicle{})

	// Add table suffix when creating tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{}, &models.Location{}, &models.Trip{}, &models.Vehicle{})
}
