package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// using asymmetric crypto/RSA keys
// location of private/public key files
const (
	privKeyPath = "keys/app.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

// private key for signing and public key for verification
var (
	verifyKey, signKey []byte
)

// read the key files before starting http handlers
func initKeys() {
	var err error

	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		panic(err)
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		panic(err)
	}
}
func GenerateJWT(name, role string) string {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	t.Claims["iss"] = "admin"
	t.Claims["UserInfo"] = struct {
		Name string
		Role string
	}{name, role}

	// set the expire time
	t.Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		panic(err)
	}
	return tokenString
}

// Middleware for validate JWT tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// validate the token
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {

		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				w.WriteHeader(http.StatusUnauthorized)
				response := Response{"Token Expired, get a new one"}
				jsonResponse(response, w)
				return

			default:
				response := Response{"Error while Parsing Token!"}
				jsonResponse(response, w)
				log.Printf("ValidationError error: %+v\n", vErr.Errors)
				return
			}

		default: // something else went wrong
			response := Response{"Error while Parsing Token!"}
			jsonResponse(response, w)
			log.Printf("Token parse error: %v\n", err)
			return
		}

	}
	if token.Valid {
		next(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		response := Response{"Invalid token"}
		jsonResponse(response, w)
	}
}

type Response struct {
	Data string `json:"error"`
}

func jsonResponse(response Response, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
