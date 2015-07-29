package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	_ "github.com/shijuvar/go-web/taskmanager/common"
	"github.com/shijuvar/go-web/taskmanager/routers"
)

//Entry point of the program
func main() {
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	//http.ListenAndServe(":5000", n)

	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
