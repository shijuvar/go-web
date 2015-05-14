package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
)

func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token := r.Header.Get("X-AppToken")
	if token == "bXlVc2VybmFtZTpteVBhc3N3b3Jk" {
		log.Printf("Authorized to the system")
		context.Set(r, "user", "Shiju Varghese")
		//Provide an access token to response's http header
		w.Header().Add("Access_Token", "6ba7b814-9dad-11d1-80b4-00c04fd430c8")
		next(w, r)
	} else {
		http.Error(w, "Not Authorized", 401)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user")
	fmt.Fprintf(w, "Welcome %s!", user)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(Authorize))
	n.UseHandler(mux)
	n.Run(":8080")
}
