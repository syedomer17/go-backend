package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"task-manager/internal/config"
)

func main() {
	router := gin.Default()

	cfg, err := config.Load()

	if err != nil {
		log.Fatal("Failed to load configuration: " + err.Error())
	}

	db, err := config.ConnectDB(cfg.MongoURI)

	if err != nil {
		log.Fatal("Failed to connect to MongoDB: " + err.Error())
	}

	_ = db

	PORT := cfg.PORT

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running",
		})
	})
	router.Run(":" + strconv.Itoa(PORT))
}
