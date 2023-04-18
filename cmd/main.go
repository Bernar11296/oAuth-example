package main

import (
	"log"
	"net/http"

	"github.com/Bernar11296/oAuth-example/handlers"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: handlers.New(),
	}
	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
