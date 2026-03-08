package controllers

import (
	"context"
	"net/http"
	"time"

	"task-manager/internal/models"
	"task-manager/internal/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	collections := models.UserCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var existingUser models.User

	err := collections.FindOne(ctx, bson.M{"email": input.Email}).Decode(&existingUser)

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email already registered",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	user := models.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	result, err := collections.InsertOne(ctx, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create user",
		})
		return
	}

	userID := result.InsertedID.(primitive.ObjectID).Hex()

	accessToken, _ := utils.GenerateAccessToken(userID)
	refreshToken, _ := utils.GenerateRefreshToken(userID)

	c.SetCookie(
		"accessToken",
		accessToken,
		86400,
		"/",
		"",
		false,
		true,
	)

	c.SetCookie(
		"refreshToken",
		refreshToken,
		604800,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "user register successfully",
	})
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	collection := models.UserCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User

	err := collection.FindOne(ctx, bson.M{
		"email": input.Email,
	}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Credentails",
		})
		return
	}

	valid := utils.CheckPassword(input.Password, user.Password)

	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Credentails",
		})
		return
	}

	userID := user.ID.Hex()

	accessToken, _ := utils.GenerateAccessToken(userID)
	refreshToken, _ := utils.GenerateRefreshToken(userID)

	c.SetCookie(
		"accessToken",
		accessToken,
		86400,
		"/",
		"",
		false,
		true,
	)

	c.SetCookie(
		"refreshToken",
		refreshToken,
		604800,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfull",
	})
}

func RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")

	if err != nil || refreshToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	claims, err := utils.VerifyToken(refreshToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Refresh token",
		})
		return
	}

	userID := claims.UserID

	newAccessToken, err := utils.GenerateAccessToken(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.SetCookie(
		"accessToken",
		newAccessToken,
		86400,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "access Token Refreshed",
	})
}

func Logout(c *gin.Context) {
	c.SetCookie(
		"accessToken",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	c.SetCookie(
		"refreshToken",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout Successfull",
	})
}