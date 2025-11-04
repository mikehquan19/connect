package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

/*
DONT CHANGE THIS FILE UNLESS YOU KNOW WHAT YOU'RE DOING

-- Hints
omitempty: omit the field if its value is empty

If you change this struct, make sure to update the database migration
scripts accordingly.(Seed, and any post requests you make)
*/

type User struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Username  string               `bson:"username" json:"username"`
	Email     string               `bson:"email" json:"email"`
	Role      string               `bson:"role" json:"role"`
	Bio       string               `bson:"bio,omitempty" json:"bio,omitempty"`
	Followers []primitive.ObjectID `bson:"followers,omitempty" json:"followers,omitempty"`
	Following []primitive.ObjectID `bson:"following,omitempty" json:"following,omitempty"`
	CreatedAt int64                `bson:"createdAt" json:"createdAt"`
}
