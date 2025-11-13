package graph

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/mikehquan19/connect/graph/model"
	"github.com/mikehquan19/connect/schema"
	"go.mongodb.org/mongo-driver/bson"
)

/*
THIS WONT SUFFER FROM THE N + 1 QUERY PROBLEM BECAUSE WE HAVE A DATALOADER
*/

// Users is the resolver for the users field.
// Query the list of users
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	// Query the list of artists (ruling out the users)
	cursor, err := r.UserCollection.Find(context.TODO(), bson.M{"role": "ARTIST"})
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

// Artworks is the resolver for the artworks field.
// Query the list of artworks embedded in the author
func (r *userResolver) Artworks(ctx context.Context, user *model.User) ([]*model.Artwork, error) {
	// Get the loaders from the request context using a key
	loaders := ctx.Value(LoadersKey).(*Loaders)

	// Load the key, in this case the user Id to the batch function
	thunk := loaders.Artworks.Load(ctx, dataloader.StringKey(user.ID))

	// Execute the batch function
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.([]*model.Artwork), nil
}
