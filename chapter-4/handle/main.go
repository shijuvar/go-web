package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})
}
func main() {
	mux := http.NewServeMux()

	mux.Handle("/welcome", messageHandler("Welcome to Go Web Development"))
	mux.Handle("/message", messageHandler("net/http is awesome"))

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
