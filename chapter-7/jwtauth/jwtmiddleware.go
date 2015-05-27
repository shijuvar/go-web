package main

// using asymmetric crypto/RSA keys

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// location of the files used for signing and verification
const (
	privKeyPath = "keys/demo.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "keys/demo.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

// keys are held in global variables
var (
	verifyKey, signKey []byte
)

// read the key files before starting http handlers
func init() {
	var err error

	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}
}

// reads the login credentials, checks them and creates the JWT token
func loginHandler(w http.ResponseWriter, r *http.Request) {

	var user User
	//decode into User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error in request body")
		return
	}
	// validate user credentials
	if user.UserName != "shijuvar" && user.Password != "pass" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Wrong info")
		return
	}

	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	t.Claims["iss"] = "admin"
	t.Claims["CustomUserInfo"] = struct {
		Name string
		Role string
	}{user.UserName, "Member"}

	// set the expire time
	t.Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Sorry, error while Signing Token!")
		log.Printf("Token Signing error: %v\n", err)
		return
	}
	response := Token{tokenString}
	jsonResponse(response, w)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{"Welcome to Admin Area"}
	jsonResponse(response, w)
}
func authMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// validate the token
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err == nil && token.Valid {
		next(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Authentication failed")
	}
}

type Response struct {
	Text string `json:"text"`
}
type Token struct {
	Token string `json:"token"`
}

func jsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

//Entry point of the program
func main() {

	r := mux.NewRouter()
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		},
	})
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.Handle("/admin", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(adminHandler)),
	))
	http.Handle("/", r)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
