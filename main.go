package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type formdetails struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	r := chi.NewRouter()

	accountDetails := formdetails{
		Name:     "Thabang",
		Password: "something",
	}
	var details []byte
	details, err := json.Marshal(accountDetails)
	if err != nil {
		log.Fatal(err)
	}

	r.Use(middleware.Logger)
	r.Get("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write(details)
	})
	fmt.Println("Listening on port: 3000")
	http.ListenAndServe(":3000", r)
}
