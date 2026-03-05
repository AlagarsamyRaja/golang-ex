package pkg

import (
	"fmt"
	"gorm/config"
	"gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(connectionString string) (*gorm.DB, error) {
	dsn, err := config.Configuration()

	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to Connect with database")
	}

	db.AutoMigrate(&models.Users{}, &models.UserRoles{}, &models.UserRoleMapping{})

	return db, nil
}
