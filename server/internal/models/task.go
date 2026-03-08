package models

import (
	"time"

	"task-manager/internal/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	Completed bool `bson:"completed" json:"completed"`
	UserID primitive.ObjectID `bson:"userId" json:"userId"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

func TaskCollection() *mongo.Collection {
	return config.DB.Database("task-manager").Collection("task")
}