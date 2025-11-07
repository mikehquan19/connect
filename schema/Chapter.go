package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chapter struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"title"`
	ArtWork    primitive.ObjectID `bson:"artWork"`
	Gallery    []string           `bson:"gallery"` // S3 keys
	Content    string             `bson:"content"`
	ChapterNum int                `bson:"chapterNum"`
	CreatedAt  primitive.DateTime `bson:"createdAt"`
	UpdatedAt  primitive.DateTime `bson:"updatedAt"`
	Meta       []any              `bson:"meta"`
}
