package main

import (
	"log"

	"github.com/josephnp732/Stocks-GraphQL/middleware"

	"github.com/joho/godotenv"
	"github.com/josephnp732/Stocks-GraphQL/http"

	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

func init() {

	//Load from local .env file (for dev ENV)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	server := gin.Default()

	middleware.SetMiddlewares(server)

	authorized := server.Group("/")
	authorized.Use(http.AuthorizeRequest())
	{
		server.GET("/", http.HandleGoogleLogin)
		server.GET("/callback", http.HandleGoogleCallback)
		authorized.GET("/playground", http.PlaygroundHandler())
		authorized.POST("/graphQL", http.GraphQLHandler())
	}

	server.Run(defaultPort)
}
