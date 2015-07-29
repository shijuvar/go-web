package routers

import (
	"github.com/gorilla/mux"
	"github.com/shijuvar/go-web/taskmanager/controllers"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/user/register", controllers.Register).Methods("POST")
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
	return router
}
