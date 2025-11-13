package graph

import (
	"context"

	"github.com/mikehquan19/connect/graph/model"
	"github.com/mikehquan19/connect/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Users is the resolver for the users field.
// Query the list of users
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	// Query the list of artists (ruling out the users)
	cursor, err := r.UserCollection.Find(
		context.TODO(),
		bson.M{"role": "ARTIST"},
	)
	if err != nil {
		return nil, err
	}

	// Unmarshal the slice of datbase users from the database
	var dbUsers []schema.User
	if err = cursor.All(context.TODO(), &dbUsers); err != nil {
		return nil, err
	}

	// Transform DB users to the GraphQL Users
	// If the client requests the artworks, then user resolver will handle that
	var gqlUsers []*model.User
	for _, dbUser := range dbUsers {
		gqlUsers = append(gqlUsers, transformUser(dbUser))
	}
	return gqlUsers, nil
}

// User is the resolver for the user field
// Query the detail of the user
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var dbUser schema.User
	err = r.UserCollection.FindOne(
		context.TODO(), bson.M{"_id": userID}).Decode(&dbUser)
	if err != nil {
		return nil, err
	}
	return transformUser(dbUser), nil
}

// Artworks is the resolver for the artworks field.
// Query the list of artworks embedded in the author
func (r *userResolver) Artworks(ctx context.Context, user *model.User) ([]*model.Artwork, error) {
	userID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return nil, err
	}
	// Checkout the latest work of the authors
	cursor, err := r.ArtCollection.Find(
		context.TODO(),
		bson.M{"author": userID},
		options.Find().SetSort(bson.D{{Key: "updatedAt", Value: -1}}),
	)
	if err != nil {
		return nil, err
	}

	// Unmarshal the Mongo data directly to GraphQL
	artworks, err := unmarshalArtworks(cursor)
	return artworks, err
}
