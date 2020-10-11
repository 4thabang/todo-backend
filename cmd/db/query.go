package db

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Brr  bool   `json:"brr"`
}

var persons = []byte(`[
	{"name": "Thabang", "age": "20", "brr": "true"},
	{"name": "Divine", "age": "2", "brr": "true"}
]`)

// QueryDB is where we read from our DB
func QueryDB() {
	var people []person
	err := json.Unmarshal(persons, &people)
	if err != nil {
		panic(err)
	}
	fmt.Println(people)
}
