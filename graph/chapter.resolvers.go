package graph

import (
	"context"

	"github.com/mikehquan19/connect/graph/model"
	"github.com/mikehquan19/connect/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Chapter is the resolver for the chapter field.
// Query a single chapter
func (r *queryResolver) Chapter(ctx context.Context, id string) (*model.Chapter, error) {
	chapID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var chap schema.Chapter
	err = r.ChapCollection.FindOne(
		context.TODO(), bson.M{"_id": chapID}).Decode(&chap)
	if err != nil {
		return nil, err
	}

	return TransformChapter(chap), nil
}

// Query the artwork of the chapter
func (r *chapterResolver) Artwork(
	ctx context.Context, chap *model.Chapter) (*model.Artwork, error) {
	artworkID, err := primitive.ObjectIDFromHex(chap.Artwork.ID)
	if err != nil {
		return nil, err
	}

	var artwork schema.Artwork
	err = r.ArtCollection.FindOne(
		context.TODO(), bson.M{"_id": artworkID}).Decode(&artwork)
	if err != nil {
		return nil, err
	}

	return transformArtwork(artwork), nil
}
