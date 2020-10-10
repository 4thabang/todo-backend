package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introducing golang",
		Description: "Come join us for a chance to learn how go",
	},
}

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Post("/create", createEvent)
	// router.Method

	fmt.Println("Listening on port `:3000`")

	http.ListenAndServe(":3000", router)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "kindly enter data with the event title and desc")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}
