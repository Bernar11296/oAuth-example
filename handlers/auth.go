package handlers

import "net/http"

func New() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("ui/")))

	mux.HandleFunc("/auth/google/login", oauthGoogleLogin)
	mux.HandleFunc("/auth/google/callback", oauthGoogleCallback)
	mux.HandleFunc("/auth/github/login", oauthGithubLogin)
	mux.HandleFunc("/auth/github/callback", oauthGithubCallback)
	return mux
}
