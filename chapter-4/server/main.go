package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Development")
}

func main() {

	http.HandleFunc("/welcome", messageHandler)

	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening...")
	server.ListenAndServe()
}
