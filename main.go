package main

import (
	"fmt"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Get("/", httpServe)

	http.ListenAndServe(":8080", r)
}

func httpServe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}
