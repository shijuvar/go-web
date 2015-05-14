package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
