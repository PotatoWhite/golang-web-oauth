package google

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

// user info url
const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

// configuration
var GoogleOAuthConfig = oauth2.Config{
	RedirectURL:  "http://localhost:3000/auth/google/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}
