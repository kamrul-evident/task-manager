package database

import (
	"task-manager/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB creates a new database connection
var DB *gorm.DB

func ConnectDatabase() error {
	db, err := gorm.Open(sqlite.Open("task_management.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto-migrate all models
	err = db.AutoMigrate(
		&models.User{},
		&models.Task{},
		// No need to migrate NameSlugDescriptionBaseModel directly as it's embedded
	)
	if err != nil {
		return err
	}

	DB = db
	return nil
}