package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shijuvar/go-web/taskmanager/common"
	"github.com/shijuvar/go-web/taskmanager/data"
	"github.com/shijuvar/go-web/taskmanager/models"
)

// Handler for HTTP Post - "/users/register"
// Add a new User document
func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		panic(err)
	}
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	c.RemoveAll(nil)
	repo := &data.UserRepository{c}
	// Insert User document
	repo.CreateUser(user)
	// Clean-up the hashpassword to eliminate it from response JSON
	user.HashPassword = nil
	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
		log.Fatal(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}

}

// Handler for HTTP Post - "/users/login"
// Authenticate with username and apssword
func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	// Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		panic(err)
	}
	loginModel := dataResource.Data
	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	// Authenticate the login user
	if user, err := repo.Login(loginUser); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	} else { //if login is successful
		w.WriteHeader(http.StatusOK)
		// Generate JWT token
		token := common.GenerateJWT(user.Email, "member")
		w.Header().Set("Content-Type", "application/json")
		// Clean-up the hashpassword to eliminate it from response JSON
		user.HashPassword = nil 
		authUser := AuthUserModel{
			User:  user,
			Token: token,
		}
		j, err := json.Marshal(AuthUserResource{Data: authUser})
		if err != nil {
			panic(err)
		}
		w.Write(j)
	}
}
