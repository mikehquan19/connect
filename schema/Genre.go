package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type Genre struct {
	ID    primitive.ObjectID `bson:"_id"`
	Genre string             `bson:"genre"`
}
