package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Bernar11296/oAuth-example/config"
	"golang.org/x/oauth2"
)

func oauthGithubLogin(w http.ResponseWriter, r *http.Request) {
	url := config.GithubOauthConfig.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func oauthGithubCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != "state" {
		fmt.Println("state is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	token, err := config.GithubOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		fmt.Printf("couldn't get token: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	client := config.GithubOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting user email: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("couldn't parse respBody: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Println(string(content))
	fmt.Fprintf(w, "Response: %s", content)
}
