package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	cors "github.com/rs/cors/wrapper/gin"

	"github.com/joho/godotenv"
	"github.com/josephnp732/Stocks-GraphQL/http"

	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

func main() {

	//Load from local .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	server := gin.Default()

	// Create new session in a cookie
	store := cookie.NewStore([]byte(http.RandToken(64)))
	store.Options(sessions.Options{
		Path: "/",
	})
	server.Use(sessions.Sessions("graphqlSession", store))
	server.Use(cors.Default())

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
