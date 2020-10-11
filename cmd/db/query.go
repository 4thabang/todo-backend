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

var persons = []person(
	{Name: "Thabang", Age: 20, Brr: true}
)

// QueryDB is where we read from our DB
func QueryDB() {
	var people []person
	err := json.Unmarshal(persons, &people)
	if err != nil {
		panic(err)
	}
	fmt.Println(people)
}
