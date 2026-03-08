package config

import (
	"context"
	"time"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("MongoDB connection error: ", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("MongoDB ping error: ", err)
	}

	DB = client 

	log.Println("MongoDB connected Successfully")

	return client, nil
}