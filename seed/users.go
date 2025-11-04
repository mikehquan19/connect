package seed

import (
	"context"
	"time"

	"github.com/mikehquan19/connect/schema"
	"github.com/mikehquan19/connect/setup"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
This file should run the first time you start the server.go

It populates the User table
*/
func SeedUsers() {
	users := []interface{}{
		schema.User{
			ID:        primitive.NewObjectID(),
			Username:  "artbymai",
			Email:     "mai@example.com",
			Role:      "artist",
			Bio:       "Digital artist and comic creator based in Dallas 🎨",
			CreatedAt: time.Now().Unix(),
		},
		schema.User{
			ID:        primitive.NewObjectID(),
			Username:  "comicfan42",
			Email:     "reader42@gmail.com",
			Role:      "reader",
			Bio:       "Big fan of indie manga and short stories 📚",
			CreatedAt: time.Now().Unix(),
		},
	}

	setup.DB.Collection("Users").InsertMany(context.TODO(), users)
}
