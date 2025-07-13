package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Define a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic("MongoDB connection failed: " + err.Error())
	}

	// Ping to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		panic("Ping failed: " + err.Error())
	}

	fmt.Println("✅ Successfully connected to MongoDB with official driver!")
}
