package common

import (
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

// Private key for signing and public key for verification
var (
	verifyKey, signKey []byte
)

// Read the key files before starting http handlers
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

// Generate JWT token
func GenerateJWT(name, role string) string {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set claims for JWT token
	t.Claims["iss"] = "admin"
	t.Claims["UserInfo"] = struct {
		Name string
		Role string
	}{name, role}

	// set the expire time for JWT token
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

		// Verify the token with public key, which is the counter part of private key
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {

		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired: //JWT expired
				DisplayAppError(w, err, "Token expired, get a new Token", 401)
				return

			default:
				DisplayAppError(w, err, "Error while parsing Token!", 500)
				return
			}

		default:
			DisplayAppError(w, err, "Error while parsing Token!", 500)
			return
		}

	}
	if token.Valid {
		next(w, r)
	} else {
		DisplayAppError(w, err, "Invalid Token", 401)
	}
}
