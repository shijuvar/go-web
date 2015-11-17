package lib

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(repo UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userStore := repo.GetAll()
		users, err := json.Marshal(userStore)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(users)
	})
}
func CreateUser(repo UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = repo.Create(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}
func SetUserRoutes() *mux.Router {
	userRepository := NewInMemoryUserRepo()
	r := mux.NewRouter()
	r.Handle("/users", CreateUser(userRepository)).Methods("POST")
	r.Handle("/users", GetUsers(userRepository)).Methods("GET")
	return r
}
