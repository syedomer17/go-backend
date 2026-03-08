package models

import (
	"task-manager/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password,omitempty" json:"-"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

func UserCollection() *mongo.Collection {
	return config.DB.Database("task_manager").Collection("users")
}