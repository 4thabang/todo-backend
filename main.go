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

	fmt.Println("Running on Port :8080")
	http.ListenAndServe(":8080", r)
}

func httpServe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}
