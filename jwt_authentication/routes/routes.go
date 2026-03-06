package routes

import (
	"authentication/handlers"
	"authentication/middlware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authHandler *handlers.AuthHandler) {

	auth := r.Group("/auth")
	{
		auth.POST("/signup", authHandler.Signup)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout")
	}

	api := r.Group("/api")
	api.Use(middlware.AuthMiddleware())
	{
		api.GET("/profile")
	}

	admin := r.Group("/admin")
	admin.Use(middlware.AuthMiddleware(), middlware.RoleMiddleware("admin"))
	{
		admin.GET("/dashboard")
	}
}
