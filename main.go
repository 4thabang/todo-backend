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

	r.Get("/get/{id}", GetTodo)
	r.Post("/add/{id}", AddTodo)

	fmt.Println("Running on Port :8080")
	http.ListenAndServe(":8080", r)
}

// GetTodo allows us to grab todos
func GetTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

// AddTodo adds a new todo using a post request
func AddTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Add Todo")
}
