package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/justinas/alice"
)

func loggingHandler(next http.Handler) http.Handler {
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	return handlers.LoggingHandler(logFile, next)
}
func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		w,
		`<doctype html>
		<html>
			<head>
				<title>Index</title>
			</head>
			<body>
				Hello Gopher!
			</body>
		</html>`,
	)

}
func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		w,
		`<doctype html>
		<html>
			<head>
				<title>About</title>
			</head>
			<body>
				Go Web development with HTTP Middleware
			</body>
		</html>`,
	)
}
func iconHandler(w http.ResponseWriter, r *http.Request) {
}
func main() {
	http.HandleFunc("/favicon.ico", iconHandler)
	indexHandler := http.HandlerFunc(index)
	aboutHandler := http.HandlerFunc(about)
	commonHandlers := alice.New(loggingHandler, handlers.CompressHandler)
	http.Handle("/", commonHandlers.ThenFunc(indexHandler))
	http.Handle("/about", commonHandlers.ThenFunc(aboutHandler))
	server := &http.Server{
		Addr: ":8080",
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
