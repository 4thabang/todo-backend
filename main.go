package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type todos struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	http.HandleFunc("/todos", TodoServer)
	http.ListenAndServe(":8080", nil)
}

// TodoServer Req/Res
func TodoServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	get()
}

func get() {
	fmt.Println("Get Request")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		log.Fatalln(err)
		return
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}

	bodyString := string(bodyBytes)
	fmt.Println("API Res:\n" + bodyString)

	var todoStruct todos
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Res Struct %+v\n", todoStruct)
}
