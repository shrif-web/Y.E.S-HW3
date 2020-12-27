package main

import (
	"log"
	"net/http"
	"os"
	"yes-blog/graph"
	"yes-blog/graph/generated"
	controller "yes-blog/internal/controller/user"
	"yes-blog/pkg/database/mongodb"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	controller.GetUserController().SetDBDriver(mongodb.NewUserMongoDriver("yes-blog","users"))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
