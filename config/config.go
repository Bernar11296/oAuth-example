package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig *oauth2.Config
var GithubOauthConfig *oauth2.Config

func Init() {
	fmt.Println(1)
	if err := godotenv.Load("./.env"); err != nil {
		log.Print("No .env file found")
	}
	GoogleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	GithubOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL_GITHUB"),
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
}
