package main

import (
    "user_service/database"
    "user_service/models"
    "user_service/controllers"
    "github.com/gin-gonic/gin"
)

func main() {

    database.ConnectToAiven()
    // Auto-migrate the User model
    database.DB.AutoMigrate(&models.User{})

    // Set up Gin router
    r := gin.Default()

    // User routes
    r.POST("/signup", controllers.Signup)
    r.POST("/login", controllers.Login)
    r.GET("/user/:id", controllers.GetUser)
    r.GET("/users", controllers.GetAllUsers)

    // Start the server
    r.Run(":8080")
}
