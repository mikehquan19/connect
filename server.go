package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mikehquan19/connect/graph"
	"github.com/mikehquan19/connect/setup"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	// Get the port and mongo uri
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mongoUri := "mongodb+srv://mikehquan19:7IoJhrnbMNAGDEQg@worshop-cluster.6meabnl.mongodb.net/?appName=Worshop-cluster"
	if mongoUri == "" {
		panic("error getting mongo uri")
	}

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		dbName = "workshop"
	}

	// Setting up the resolver
	database := setup.ConnectDB(mongoUri, dbName)
	resolver := graph.Resolver{
		UserCollection: database.Collection("users"),
		ArtCollection:  database.Collection("artworks"),
		ChapCollection: database.Collection("chapters"),
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver}))
	srv.Use(extension.FixedComplexityLimit(100))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", resolver.Middleware()(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
