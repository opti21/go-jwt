package database

import (
	"github.com/opti21/go-jwt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := ""
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not open postgres connection")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}