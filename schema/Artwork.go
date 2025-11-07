package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

// Artwork status
type Status string

const (
	ONGOING   Status = "ONGOING"
	COMPLETED Status = "COMPLETED"
)

// Artwork type
type ArtType string

const (
	NOVEL ArtType = "NOVEL"
	MANGA ArtType = "MANGA"
)

type Artwork struct {
	ID            primitive.ObjectID   `bson:"_id"`
	Title         string               `bson:"title"`
	Summary       string               `bson:"summary"`
	Cover         string               `bson:"cover"`
	Author        primitive.ObjectID   `bson:"author"`
	Status        Status               `bson:"status"`
	ArtType       ArtType              `bson:"ArtType"`
	LatestChapter int                  `bson:"latestChapter"`
	Chapters      []primitive.ObjectID `bson:"chapters"`
	CreatedAt     primitive.DateTime   `bson:"createdAt"`
	UpdatedAt     primitive.DateTime   `bson:"updatedAt"`
	Genres        []primitive.ObjectID `bson:"genres"`
	Meta          []any                `bson:"meta"`
}
