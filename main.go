package main

import (
	"fmt"
	"os"

	"github.com/josephnp732/Stocks-GraphQL/http"

	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	server := gin.Default()

	server.GET("/", http.PlaygroundHandler())

	server.POST("/graphQL", http.GraphQLHandler())

	server.Run(defaultPort)
}
