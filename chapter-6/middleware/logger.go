package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}
func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index handler")
	fmt.Fprintf(w, "Welcome!")
}
func about(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing about handler")
	fmt.Fprintf(w, "Go Middleware")
}
func iconHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/favicon.ico", iconHandler)
	indexHandler := http.HandlerFunc(index)
	aboutHandler := http.HandlerFunc(about)
	http.Handle("/", loggingHandler(indexHandler))
	http.Handle("/about", loggingHandler(aboutHandler))
	server := &http.Server{
		Addr: ":8080",
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
