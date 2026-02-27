package main

import (
	"log"
	"user_mangement/config"
	"user_mangement/handlers"
	"user_mangement/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	cfg := config.LoadConfig()

	db := config.ConnectToDB(cfg)
	db.AutoMigrate(&models.User{})

	r := gin.Default()

	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	userHandler := handlers.NewUserHandler(db)
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetUsers)
	r.GET("/users/:id", userHandler.GetUserById)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.PATCH("/users/:id", userHandler.PatchUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	r.Run(":" + cfg.Port)
}
