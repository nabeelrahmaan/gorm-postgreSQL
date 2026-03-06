package handlers

import (
	"authentication/dto"
	"authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *service.AuthService
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var req dto.SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	err := h.Service.Signup(req.Name, req.Email, req.Password)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req *dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid credentials"})
		return
	}

	access, refresh, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	response := &dto.AuthResponse{
		AccessToken: access,
		RefreshToken: refresh,
	}


	c.JSON(200, response)
}