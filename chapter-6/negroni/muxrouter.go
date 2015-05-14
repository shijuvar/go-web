package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")
}
