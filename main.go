package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/4thabang/todo-backend/cmd/handler"
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

	router.Get("/", handler.Handler)
	router.Get("/api", fetchAPI)
	router.Post("/create", createEvent)
	router.Get("/events", getEvents)
	router.Get("/events/{id}", getOneEvent)
	router.Delete("/events/{id}", deleteEvent)
	router.Get("/ping", ping)

	fmt.Println("\nListening on port :3000")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world. Welcome home")
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pong")
}

type todos struct {
	UserID     int    `json:"userId"`
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Complteted bool   `json:"completed"`
}

func fetchAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var todo todos
	json.Unmarshal(data, &todo)
	fmt.Fprint(w, string(data))
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(events)
}

// package main

// import (
// 	"fmt"

// 	"net/http"

// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/chi/middleware"
// )

// func main() {
// 	r := chi.NewRouter()

// 	r.Use(middleware.Logger)

// 	r.Get("/get/{id}", GetTodo)
// 	r.Post("/add/{id}", AddTodo)

// 	fmt.Println("Running on Port :8080")
// 	http.ListenAndServe(":8080", r)
// }

// // GetTodo allows us to grab todos
// func GetTodo(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello world")
// }

// // AddTodo adds a new todo using a post request
// func AddTodo(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Add Todo")
// }
