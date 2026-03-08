package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"task-manager/internal/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateTaskInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func CreateTask(c *gin.Context) {
	var input CreateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userIdValue, exists := c.Get("userId")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userID, err := primitive.ObjectIDFromHex(userIdValue.(string))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid User ID",
		})
		return
	}

	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
		UserID:      userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	collection := models.TaskCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create task",
		})
		return
	}

	task.ID = result.InsertedID.(primitive.ObjectID)

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func GetTasks(c *gin.Context) {
	userIdValue, exists := c.Get("userId")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userID, err := primitive.ObjectIDFromHex(userIdValue.(string))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid User ID",
		})
		return
	}

	collection := models.TaskCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//pagination
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)

	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid page number",
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)

	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid limit number",
		})
		return
	}

	skip := (page - 1) * limit

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(skip))
	findOptions.SetSort(bson.M{"createdAt": -1})

	cursor, err := collection.Find(ctx, bson.M{
		"userId": userID,
	}, findOptions)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch tasks",
		})
		return
	}

	var tasks []models.Task

	if err = cursor.All(ctx, &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to decode tasks",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

type UpdateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   *bool  `json:"completed"`
}

func UpdateTask(c *gin.Context) {
	taskId := c.Param("id")

	objectTaskID, err := primitive.ObjectIDFromHex(taskId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Task ID",
		})
		return
	}

	userIdValue, exists := c.Get("userId")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	userID, err := primitive.ObjectIDFromHex(userIdValue.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid User ID",
		})
		return
	}

	var input UpdateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	update := bson.M{}
	setFields := bson.M{}

	if input.Title != "" {
		setFields["title"] = input.Title
	}

	if input.Description != "" {
		setFields["description"] = input.Description
	}

	if input.Completed != nil {
		setFields["completed"] = *input.Completed
	}

	setFields["updatedAt"] = time.Now()

	update["$set"] = setFields

	collection := models.TaskCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(
		ctx,
		bson.M{
			"_id":    objectTaskID,
			"userId": userID,
		},
		update,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to updated task",
		})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task updated",
	})
}

func DeleteTask(c *gin.Context) {
	taskId := c.Param("id")

	objectTaskID, err := primitive.ObjectIDFromHex(taskId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Task ID",
		})
		return
	}

	userIdValue, exists := c.Get("userId")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userID, err := primitive.ObjectIDFromHex(userIdValue.(string))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}

	collection := models.TaskCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(
		ctx,
		bson.M{
			"_id":    objectTaskID,
			"userId": userID,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete task",
		})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "task deleted",
	})
}
