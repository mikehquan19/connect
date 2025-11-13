package graph

import (
	"context"

	"github.com/mikehquan19/connect/graph/model"
	"github.com/mikehquan19/connect/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Artworks is the resolver for the artworks field.
// Query the list of latest artworks
func (r *queryResolver) Artworks(ctx context.Context) ([]*model.Artwork, error) {
	// Sort by the last updated date and then created date
	cursor, err := r.ArtCollection.Find(
		context.TODO(),
		bson.M{},
		options.Find().SetSort(bson.D{
			{Key: "updatedAt", Value: -1},
			{Key: "createdAt", Value: -1},
		}),
	)
	if err != nil {
		return nil, err
	}

	artworks, err := unmarshalArtworks(cursor)
	return artworks, err
}

// Artwork is the resolver for the artwork field.
// Query a single artwork
func (r *queryResolver) Artwork(ctx context.Context, id string) (*model.Artwork, error) {
	artworkID, err := primitive.ObjectIDFromHex(id)
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

// Author is the resolver for the users field.
// Query the author embedded in the artwork
func (r *artworkResolver) Author(ctx context.Context, work *model.Artwork) (*model.User, error) {
	authorId, err := primitive.ObjectIDFromHex(work.Author.ID)
	if err != nil {
		return nil, err
	}

	var author schema.User
	err = r.UserCollection.FindOne(
		context.TODO(), bson.M{"_id": authorId}).Decode(&author)
	if err != nil {
		return nil, err
	}

	return transformUser(author), nil
}

// Chapters is the resolver for the chapters field.
// Query the list of chapters of the artwork
func (r *artworkResolver) Chapters(ctx context.Context, work *model.Artwork) ([]*model.Chapter, error) {
	convertedID, err := primitive.ObjectIDFromHex(work.ID)
	if err != nil {
		return nil, err
	}

	cursor, err := r.ChapCollection.Find(
		context.TODO(),
		bson.M{"artWork": convertedID},
		options.Find().SetSort(bson.D{{Key: "createdAt", Value: 1}}),
	)
	if err != nil {
		return nil, err
	}

	chapters, err := unmarshalChapters(cursor)
	return chapters, err
}
