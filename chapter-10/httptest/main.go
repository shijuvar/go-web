package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:email"`
}

var userStore = []User{}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := json.Marshal(userStore)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(users)

}
func createUser(w http.ResponseWriter, r *http.Request) {

	var user User
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Validate the User entity
	err = validate(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Insert User entity into User Store
	userStore = append(userStore, user)
	w.WriteHeader(http.StatusCreated)
}

// Validate User entity
func validate(user User) error {
	for _, u := range userStore {
		if u.Email == user.Email {
			return errors.New("The Email is already exists")
		}
	}
	return nil
}
func SetUserRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users", getUsers).Methods("GET")
	return r
}

func main() {
	http.ListenAndServe(":8080", SetUserRoutes())
}
