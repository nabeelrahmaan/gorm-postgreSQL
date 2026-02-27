package handlers

import (
	"log"
	"net/http"
	"user_mangement/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	if err := h.DB.Create(&user).Error; err != nil {
		log.Println("Create user error:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	var users []models.User

	if err := h.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found",})
	}

	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON",})
		return
	}

	user.Name = input.Name
	user.Email = input.Email

	if err := h.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user",})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) PatchUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found",})
		return
	}

	var input map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON",})
		return
	}

	delete(input, "id")
	delete(input, "created_at")

	if err := h.DB.Model(&user).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user",})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found",})
		return
	}

	if err := h.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user",})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted succesfully",})
}
