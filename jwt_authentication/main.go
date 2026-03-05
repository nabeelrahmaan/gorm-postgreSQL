package main

import (
	"authentication/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env files")
	}

	cfg := config.LoadConfig()
	config.ConnectToDb(cfg)

	router := gin.Default()

	router.Run(":" + cfg.Port)
}