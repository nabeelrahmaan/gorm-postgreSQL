package config

import (
	"authentication/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb(cfg *Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	database, err := gorm.Open(postgres.Open(dsn),&gorm.Config{}) 
	if err != nil {
		log.Fatalf("Failed to connect to database")
	}

	database.AutoMigrate(&models.User{}, &models.RefreshToken{} )

	return database
}