package main

import (
	"app/graph"
	"app/graph/generated"
  "github.com/99designs/gqlgen/handler"
  "github.com/gin-gonic/gin"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
  h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

  return func(c *gin.Context) {
    h.ServeHTTP(c.Writer, c.Request)
  }
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
  h := handler.Playground("GraphQL", "/query")

  return func(c *gin.Context) {
    h.ServeHTTP(c.Writer, c.Request)
  }
}

func main() {
  // Setting up Gin
  r := gin.Default()
  r.POST("/query", graphqlHandler())
  r.GET("/", playgroundHandler())
    r.Run()
}
