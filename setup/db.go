package setup

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB connects the server with the Mongo Database
func ConnectDB(mongoUri string, databaseName string) *mongo.Database {
	var database *mongo.Database

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	database = client.Database(databaseName)
	log.Println("Connected to MongoDB:", databaseName)
	return database
}
