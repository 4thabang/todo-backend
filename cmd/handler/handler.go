package handler

import (
	"fmt"
	"net/http"
)

// Handler allows us to play with our api
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
