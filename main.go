package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/4thabang/todo-backend/cmd/db"
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

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/read", db.ReadData)

	router.Get("/", home)
	router.Post("/create", createEvent)
	router.Get("/events", getEvents)
	router.Get("/events/{id}", getOneEvent)
	router.Delete("/events/{id}", deleteEvent)
	router.Get("/ping", ping)

	fmt.Println("\nListening on port :8080")

	http.ListenAndServe(":8080", router)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world. Welcome home")
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pong")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "kindly enter data with the event title and desc")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)

	w.WriteHeader(http.StatusCreated)

	if http.StatusOK == 200 {
		fmt.Fprintln(w, "Post has been added!")
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := chi.URLParam(r, "id")

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The event with ID: %v.\nHas been deleted successfully", eventID)
		}
	}
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	eventID := chi.URLParam(r, "id")

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			err := json.NewEncoder(w).Encode(singleEvent)
			if err != nil {
				panic(err)
			}
		}
	}
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
