package main

import (
	"authentication/config"
	"authentication/handlers"
	"authentication/repository"
	"authentication/routes"
	"authentication/service"
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
	db := config.ConnectToDb(cfg)

	repository := &repository.UserRepository{DB: db,}

	service := &service.AuthService{Repo: repository,}

	handler := &handlers.AuthHandler{Service: service,}


	r := gin.Default()
	routes.RegisterRoutes(r, handler)

	r.Run(":" + cfg.Port)
}