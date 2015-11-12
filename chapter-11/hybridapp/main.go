// +build !appengine

package main

import (
	"net/http"

	_ "github.com/shijuvar/go-web/chapter-11/hybridapplib"
)

func main() {
	http.ListenAndServe("localhost:8080", nil)
}
