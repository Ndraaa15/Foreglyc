package oauth

import (
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func New() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     viper.GetString("oauth.google.client_id"),
		ClientSecret: viper.GetString("oauth.google.client_secret"),
		RedirectURL:  "http://localhost:8000/auth/callback",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	return conf
}
