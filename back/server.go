package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"yes-blog/graph"
	"yes-blog/graph/generated"
	postController "yes-blog/internal/controller/post"
	userController "yes-blog/internal/controller/user"
	"yes-blog/internal/middleware/auth"
	"yes-blog/internal/middleware/ggcontext"
	"yes-blog/pkg/database/mongodb"
)

const defaultPort = "8080"
const queryComplexity = 8


func main() {

	//setting a mongodb driver for DBDriver filed of our controllers instance
	userController.GetUserController().SetDBDriver(mongodb.NewUserMongoDriver("yes-blog", "users"))
	postController.GetPostController().SetDBDriver(mongodb.NewPostMongoDriver("yes-blog", "users"))

	// Setting up Gin
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(ggcontext.GinContextToContextMiddleware())
	r.Use(auth.Middleware())

	// routing
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	//let it begin
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	r.Run(":"+defaultPort)

}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	//srv.Use(extension.FixedComplexityLimit(queryComplexity))

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	pg := playground.Handler("Yes-blog playground", "/query")

	return func(c *gin.Context) {
		pg.ServeHTTP(c.Writer, c.Request)
	}
}