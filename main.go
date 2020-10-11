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
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introducing golang",
		Description: "Come join us for a chance to learn how go",
	},
	{
		ID:          "2",
		Title:       "This is a rest api",
		Description: "Rest apis are cool and should be used often",
	},
}

var (
	router = chi.NewRouter()
)

func main() {
	router.Use(middleware.Logger)

	router.Post("/create", createEvent)
	router.Get("/events", getEvents)
	router.Get("/events/{id}", getOneEvent)

	fmt.Println("Listening on port :8080")

	http.ListenAndServe(":8080", router)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	// Gets 'event' and uses its type structure
	var newEvent event
	// Reads the body entry
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "kindly enter data with the event title and desc")
	}

	// Takes body entry value and passes it into `newEvent` in it's set structure
	json.Unmarshal(reqBody, &newEvent)
	// We simply append the 'newEvent' entry into the 'events' variable (database)
	events = append(events, newEvent)
	// HTTP status code in header
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	eventID := chi.URLParam(r, "id")

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
