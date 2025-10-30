package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    
    "repsheets-go/handlers"
)

// SetupRoutes configureert alle routes
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
    // UserHandler aanmaken
    userHandler := &handlers.UserHandler{
        DB: db,
    }
    
    // User routes
    userRoutes := router.Group("/users")
    {
        userRoutes.GET("/", userHandler.GetUsers)      // GET /users
        userRoutes.POST("/", userHandler.CreateUser)   // POST /users
    }
}