package main

import (
	"fmt"
	"io/ioutil"
	"log"

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

type todoStr struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func httpServe(w http.ResponseWriter, r *http.Request) {
	// Set HTTP Headers for JSON and CORS access
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create a HTTP GET request and handle err
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}

	// Close the body last w/ 'defer'
	defer res.Body.Close()

	// Reads data from body and returns w/ err handler
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print read values in string JSON format
	fmt.Println(string(data))
	fmt.Fprintln(w, string(data))
}
