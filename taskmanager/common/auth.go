package common

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/context"
)

// AppClaims provides custom claim for JWT
type AppClaims struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// using asymmetric crypto/RSA keys
// location of private/public key files
const (
	// openssl genrsa -out app.rsa 1024
	privKeyPath = "keys/tm.rsa"
	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "keys/tm.rsa.pub"
)

// Private key for signing and public key for verification
var (
	//verifyKey, signKey []byte
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// Read the key files before starting http handlers
func initKeys() {

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	/*
		signKey, err = ioutil.ReadFile(privKeyPath)
		if err != nil {
			log.Fatalf("[initKeys]: %s\n", err)
		}

		verifyKey, err = ioutil.ReadFile(pubKeyPath)
		if err != nil {
			log.Fatalf("[initKeys]: %s\n", err)
			panic(err)
		}
	*/
}

// GenerateJWT returns a JWT token
func GenerateJWT(name, role string) (string, error) {
	// Create the Claims
	claims := AppClaims{
		name,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			Issuer:    "admin",
		},
	}
	log.Println("before token")
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	log.Println("before key")
	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	log.Println("after key")
	return ss, nil
	/*
		// create a signer for rsa 256
		t := jwt.New(jwt.GetSigningMethod("RS256"))
		t1 := jwt.NewWithClaims()

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
			return "", err
		}
		return tokenString, nil
	*/
}

// Authorize Middleware for validating JWT tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// Get token from request
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})
	/*
		// validate the token
		token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {

			// Verify the token with public key, which is the counter part of private key
			return verifyKey, nil
		})
	*/

	if err != nil {
		switch err.(type) {

		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired: //JWT expired
				DisplayAppError(
					w,
					err,
					"Access Token is expired, get a new Token",
					401,
				)
				return

			default:
				DisplayAppError(w,
					err,
					"Error while parsing the Access Token!",
					500,
				)
				return
			}

		default:
			DisplayAppError(w,
				err,
				"Error while parsing Access Token!",
				500)
			return
		}

	}
	if token.Valid {
		context.Set(r, "user", token.Claims.(*AppClaims).UserName)
		next(w, r)
	} else {
		DisplayAppError(
			w,
			err,
			"Invalid Access Token",
			401,
		)
	}
}

// TokenFromAuthHeader is a "TokenExtractor" that takes a given request and extracts
// the JWT token from the Authorization header.
func TokenFromAuthHeader(r *http.Request) (string, error) {
	// Look for an Authorization header
	if ah := r.Header.Get("Authorization"); ah != "" {
		// Should be a bearer token
		if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
			return ah[7:], nil
		}
	}
	return "", errors.New("No token in the HTTP request")
}
