package http

import (
	"log"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/josephnp732/Stocks-GraphQL/apq"
	"github.com/josephnp732/Stocks-GraphQL/graph"
	"github.com/josephnp732/Stocks-GraphQL/graph/generated"
)

// PlaygroundHandler returns the handler for the GraphQL playground
func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/graphQL")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// GraphQLHandler return sthe handler for the graphQL URL
func GraphQLHandler() gin.HandlerFunc {

	cache, err := apq.NewCache(24 * time.Hour)
	if err != nil {
		log.Fatalf("cannot create APQ redis cache: %v", err)
	}

	c := generated.Config{Resolvers: &graph.Resolver{}}
	srv := handler.GraphQL(
		generated.NewExecutableSchema(c),
		handler.EnablePersistedQueryCache(cache),
	)

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
