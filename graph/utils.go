package graph

import (
	"context"
	"time"

	"github.com/mikehquan19/connect/graph/model"
	"github.com/mikehquan19/connect/schema"
	"go.mongodb.org/mongo-driver/mongo"
)

// TransformUser transforms a MongoDB user into a GraphQL user
func transformUser(dbUser schema.User) *model.User {
	return &model.User{
		ID:        dbUser.ID.Hex(),
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		Role:      model.Role(dbUser.Role),
		Bio:       dbUser.Bio,
		ImageURI:  dbUser.ImageUri,
		CreatedAt: dbUser.CreatedAt.Time().Format(time.RFC3339),
	}
}

// TransformArtwork transforms a MongoDB arkwork into a GraphQL artwork
func transformArtwork(dbArtwork schema.Artwork) *model.Artwork {
	return &model.Artwork{
		ID:            dbArtwork.ID.Hex(),
		Title:         dbArtwork.Title,
		Summary:       dbArtwork.Summary,
		Cover:         dbArtwork.Cover,
		Author:        &model.User{ID: dbArtwork.Author.Hex()}, // Needed to get the query
		Status:        model.Status(dbArtwork.Status),
		ArtType:       model.ArtType(dbArtwork.ArtType),
		LatestChapter: int32(dbArtwork.LatestChapter),
		CreatedAt:     dbArtwork.CreatedAt.Time().Format(time.RFC3339),
		UpdatedAt:     dbArtwork.UpdatedAt.Time().Format(time.RFC3339),
	}
}

// unmarshalArtworks marshals the data from mongo to artworks
func unmarshalArtworks(mongoCursor *mongo.Cursor) ([]*model.Artwork, error) {
	// Unmarshal the slice of artworks
	var dbArtworks []schema.Artwork
	err := mongoCursor.All(context.TODO(), &dbArtworks)
	if err != nil {
		return nil, err
	}

	// Transform database artworks to graphql artworks
	var gqlArtworks []*model.Artwork
	for _, dbArtwork := range dbArtworks {
		gqlArtworks = append(gqlArtworks, transformArtwork(dbArtwork))
	}
	return gqlArtworks, nil
}

// TransformChapter transforms a MongoDB chapter into a GraphQL chapter
func TransformChapter(dbChapter schema.Chapter) *model.Chapter {
	return &model.Chapter{
		ID:         dbChapter.ID.Hex(),
		Title:      dbChapter.Title,
		Artwork:    &model.Artwork{ID: dbChapter.ID.Hex()},
		Gallery:    dbChapter.Gallery,
		Content:    &dbChapter.Content,
		ChapterNum: int32(dbChapter.ChapterNum),
		CreatedAt:  dbChapter.CreatedAt.Time().Format(time.RFC3339),
		UpdatedAt:  dbChapter.UpdatedAt.Time().Format(time.RFC3339),
	}
}

// unmarshalChapters marshals the data from mongo to chapters
// and then transform them to GraphQL responses.
func unmarshalChapters(mongoCursor *mongo.Cursor) ([]*model.Chapter, error) {
	// Unmarshal the slice of chapters
	var dbChapters []schema.Chapter
	err := mongoCursor.All(context.TODO(), &dbChapters)
	if err != nil {
		return nil, err
	}

	// Transform database chapters to graphql chapters
	var gqlChapters []*model.Chapter
	for _, dbChapter := range dbChapters {
		gqlChapters = append(gqlChapters, TransformChapter(dbChapter))
	}
	return gqlChapters, nil
}
