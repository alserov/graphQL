package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alserov/graphQL/db"
	"github.com/alserov/graphQL/graph"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dts, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	defer dts.Close()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
