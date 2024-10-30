package controllers

import (
	"net/http"
	"user_service/models"
	"user_service/services"

	"github.com/gin-gonic/gin"
)

// Signup handles user registration
func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	createdUser, err := services.Signup(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Exclude password field in response
	createdUser.Password = ""
	c.JSON(http.StatusCreated, createdUser)
}

// Login handles user authentication
func Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user, err := services.Login(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Exclude password field in response
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

// GetUser fetches a user by id
func GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := services.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Exclude password field in response
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

// GetAllUsers fetches all registered users
func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	// Exclude password field in the response
	for i := range users {
		users[i].Password = ""
	}

	c.JSON(http.StatusOK, users)
}
