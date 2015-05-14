package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index handler")
	fmt.Fprintf(w, "Welcome!")
}
func about(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing about handler")
	fmt.Fprintf(w, "Go Middleware")
}
func iconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./favicon.ico")
}
func main() {
	http.HandleFunc("/favicon.ico", iconHandler)
	indexHandler := http.HandlerFunc(index)
	aboutHandler := http.HandlerFunc(about)
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	http.Handle("/", handlers.LoggingHandler(logFile, handlers.CompressHandler(indexHandler)))
	http.Handle("/about", handlers.LoggingHandler(logFile, handlers.CompressHandler(aboutHandler)))
	server := &http.Server{
		Addr: ":8080",
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
