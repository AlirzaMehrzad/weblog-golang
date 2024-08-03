package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collections struct {
	Users    *mongo.Collection
	Posts    *mongo.Collection
	Comments *mongo.Collection
	Files    *mongo.Collection
}

// InitDB initializes the MongoDB connection and returns the client and collections
func InitDB() (*mongo.Client, *Collections) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	// Database and collections
	db := client.Database("goApp-login-register")
	collections := &Collections{
		Users:    db.Collection("users"),
		Posts:    db.Collection("posts"),
		Comments: db.Collection("comments"),
		Files:    db.Collection("files"),
	}

	return client, collections
}
