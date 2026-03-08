package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"task-manager/internal/config"
	"task-manager/internal/middleware"
	"task-manager/internal/routes"
	"task-manager/pkg/logger"
)

func main() {

	logger.Init()
	defer logger.Log.Sync()

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

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
		},
		AllowCredentials: true,
	}))

	router.Use(middleware.RequestLogger())
	router.Use(middleware.RateLimiter())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running",
		})
	})

	routes.SetupRoutes(router)

	router.Run(":" + strconv.Itoa(PORT))
}
