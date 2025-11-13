package graph

import (
	"context"
	"net/http"

	"github.com/graph-gophers/dataloader"
	"github.com/mikehquan19/connect/graph/model"
	"github.com/mikehquan19/connect/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Loaders struct {
	Artworks *dataloader.Loader
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
				Artworks: dataloader.NewBatchedLoader(batchArtworks(e.ArtCollection)),
			}
			// Think of it context as a map, here we put new pair in
			ctx := context.WithValue(r.Context(), LoadersKey, loaders)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// batchArtworks batches all the artwork resolvers and get
// all the artworks with one database query
func batchArtworks(collection *mongo.Collection) dataloader.BatchFunc {
	return func(ctx context.Context, k dataloader.Keys) []*dataloader.Result {
		// Convert all the IDs
		var IDs []primitive.ObjectID
		for _, key := range k {
			ID, err := primitive.ObjectIDFromHex(key.String())
			if err != nil {
				return []*dataloader.Result{{Error: err}}
			}
			IDs = append(IDs, ID)
		}

		// One query database
		// Unmarshal list of artworks from the database
		cursor, err := collection.Find(ctx, bson.M{"author": bson.M{"$in": IDs}})
		if err != nil {
			return []*dataloader.Result{{Error: err}}
		}
		var dbArtworks []schema.Artwork
		err = cursor.All(ctx, &dbArtworks)
		if err != nil {
			return []*dataloader.Result{{Error: err}}
		}

		authorToArtwork := make(map[string][]*model.Artwork)
		for _, artwork := range dbArtworks {
			ID := artwork.Author.Hex()
			authorToArtwork[ID] = append(authorToArtwork[ID], transformArtwork(artwork))
		}

		var results []*dataloader.Result
		for _, key := range k {
			results = append(results, &dataloader.Result{
				Data: authorToArtwork[key.String()],
			})
		}

		return results
	}
}
