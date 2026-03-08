package controllers

import (
	"context"
	"net/http"
	"time"

	"task-manager/internal/models"
	"task-manager/internal/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type RegisterInput struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func Register(c *gin.Context){
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	collections := models.UserCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

	defer cancel()

	var existingUser models.User 

	err := collections.FindOne(ctx, bson.M{"email": input.Email}).Decode(&existingUser)
	
	if err == nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "email already registered",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)

	if err != nil {
		c.JSON(http.statusInternalServerError,gin.H{
			"error": "failed to hash password",
		})
		return
	}

	user := models.User{
		Name: input.Name,
		Email: input.Email,
		Password: hashedPassword,
		CreatedAt: time.Now(),
	}
}