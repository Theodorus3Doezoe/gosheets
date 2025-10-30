package handlers

import (
    "net/http"
    //"strconv"
    "golang.org/x/crypto/bcrypt"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    
    "repsheets-go/models"
	"repsheets-go/schemas"
)

type UserHandler struct {
    DB *gorm.DB
}

// GetUsers haalt alle users op uit de database
func (h *UserHandler) GetUsers(c *gin.Context) {
    var users []models.User
    
    result := h.DB.Find(&users)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        return
    }
    
    c.JSON(http.StatusOK, users)
}

// CreateUser maakt een nieuwe user aan
func (h *UserHandler) CreateUser(c *gin.Context) {
    var req schemas.CreateUserRequest
    
    err := c.ShouldBindJSON(&req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
	var existingUser models.User
    checkResult := h.DB.Where("email = ?", req.Email).First(&existingUser)
    if checkResult.Error == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    
    newUser := models.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: string(hashedPassword),
    }
    
    createResult := h.DB.Create(&newUser)
    if createResult.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }
    
    response := schemas.UserResponse{
        ID:        newUser.ID,
        Name:      newUser.Name,
        Email:     newUser.Email,
        CreatedAt: newUser.CreatedAt.Format("2006-01-02 15:04:05"),
    }
    
    c.JSON(http.StatusCreated, response)
}
