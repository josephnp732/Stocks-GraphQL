package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
	oauthStateString  = RandToken(64)
)

func init() {

	callbackURL, exists := os.LookupEnv("CALLBACK_URL")
	if !exists {
		callbackURL = "http://localhost:8080/callback"
	}

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  callbackURL,
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

// HandleGoogleLogin Login Handler for OAuth
func HandleGoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)

	session := sessions.Default(c)

	// if session is still valid then authorize user
	if accessToken := session.Get("access_token"); accessToken != nil {
		http.Redirect(c.Writer, c.Request, "/playground", http.StatusTemporaryRedirect)
	}

	http.Redirect(c.Writer, c.Request, url, http.StatusTemporaryRedirect)
}

// HandleGoogleCallback is the callback function for OAuth
func HandleGoogleCallback(c *gin.Context) {
	_, err := getUserInfo(c.Request.FormValue("state"), c.Request.FormValue("code"), c)

	if err != nil {
		log.Println(err.Error())
		http.Error(c.Writer, `User Unauthorized! Please Login`, http.StatusUnauthorized)
	}

	//log.Printf("Content: %s\n", content)
	http.Redirect(c.Writer, c.Request, "/playground", http.StatusTemporaryRedirect)
}

// AuthorizeRequest is used to authorize a request for a certain end-point group.
func AuthorizeRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		v := session.Get("access_token")
		if v == nil {
			// if not authorized redirect to login page
			http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
			return
		}
		c.Next()
	}
}

func getUserInfo(state string, code string, c *gin.Context) ([]byte, error) {

	session := sessions.Default(c)

	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	// Set session key-values
	session.Set("access_token", token.AccessToken)
	session.Set("expiry", token.Expiry.String())
	session.Set("refresh_token", token.RefreshToken)

	// session age = age of access token
	session.Options(sessions.Options{
		MaxAge: int(token.Expiry.Sub(time.Now()).Seconds()),
	})
	err = session.Save()

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}
