package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/josephnp732/Stocks-GraphQL/http"
	cors "github.com/rs/cors/wrapper/gin"
)

// SetMiddlewares Sets required Middlewares
func SetMiddlewares(server *gin.Engine) {

	// Create new session in a cookie
	store := cookie.NewStore([]byte(http.RandToken(64)))
	store.Options(sessions.Options{
		Path: "/",
	})
	server.Use(sessions.Sessions("graphqlSession", store))

	// CORS middleware
	server.Use(cors.Default())
}
