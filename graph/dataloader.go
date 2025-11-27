package graph

import (
	"context"
	"net/http"

	"github.com/graph-gophers/dataloader"
	"github.com/mikehquan19/connect/graph/model"
	"github.com/mikehquan19/connect/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Loaders struct {
	Artworks *dataloader.Loader
	Author   *dataloader.Loader
}

// The key that is used to retrieve the loaders from the context
type ctxKey string

const LoadersKey ctxKey = "dataloaders"

// Middleware to inject the loaders in every incoming request context
func (e *Resolver) Middleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create the loaders with all the batch functions
			loaders := &Loaders{
				Artworks: dataloader.NewBatchedLoader(e.batchArtworks()),
				Author:   dataloader.NewBatchedLoader(e.batchAuthor()),
			}

			ctx := context.WithValue(r.Context(), LoadersKey, loaders)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// batchArtworks batches all the artwork resolvers and perform 1 query on the database
func (e *Resolver) batchArtworks() dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		userIDs, err := convertIDs(keys)
		if err != nil {
			return []*dataloader.Result{{Error: err}}
		}

		// Unmarshal list of artworks from the database using only query
		var dbArtworks []schema.Artwork
		cursor, err := e.ArtCollection.Find(ctx, bson.M{"author": bson.M{"$in": userIDs}})
		if err != nil {
			return []*dataloader.Result{{Error: err}}
		}
		err = cursor.All(ctx, &dbArtworks)
		if err != nil {
			return []*dataloader.Result{{Error: err}}
		}

		// Map from the userID to the GraphQL artworks
		userToArtwork := make(map[string][]*model.Artwork)
		for _, dbArtwork := range dbArtworks {
			id := dbArtwork.Author.Hex()
			userToArtwork[id] = append(userToArtwork[id], transformArtwork(dbArtwork))
		}

		artworkResults := make([]*dataloader.Result, len(keys))
		for i, key := range keys {
			artworkResults[i] = &dataloader.Result{
				Data: userToArtwork[key.String()],
			}
		}
		return artworkResults
	}
}

// batchAuthor batches all the userResolver
func (e *Resolver) batchAuthor() dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		artworkIDs, err := convertIDs(keys)
		if err != nil {
			return []*dataloader.Result{{Error: err}}
		}

		var dbUsers []schema.User
		cursor, err := e.UserCollection.Find(ctx, bson.M{"_id": bson.M{"$in": artworkIDs}})
		if err != nil {
			return []*dataloader.Result{{Error: err}}
		}
		err = cursor.All(ctx, &dbUsers)
		if err != nil {
			return []*dataloader.Result{{Error: err}}
		}

		artworkToAuthor := make(map[string]*model.User)
		for _, dbUser := range dbUsers {
			artworkToAuthor[dbUser.ID.Hex()] = transformUser(dbUser)
		}

		authorResults := make([]*dataloader.Result, len(keys))
		for i, key := range keys {
			authorResults[i] = &dataloader.Result{
				Data: artworkToAuthor[key.String()],
			}
		}
		return authorResults
	}
}

// convertIDs converts the list of keys to the list of Mongo IDs
func convertIDs(keys dataloader.Keys) ([]primitive.ObjectID, error) {
	convertedIDs := make([]primitive.ObjectID, len(keys))
	for i, key := range keys {
		convertedID, err := primitive.ObjectIDFromHex(key.String())
		if err != nil {
			return nil, err
		}
		convertedIDs[i] = convertedID
	}
	return convertedIDs, nil
}
