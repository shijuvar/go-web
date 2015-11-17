package main

import (
	"net/http"

	"github.com/shijuvar/go-web/chapter-10/httptestbdd/lib"
)

func main() {
	routers := lib.SetUserRoutes()
	http.ListenAndServe(":8080", routers)
}
