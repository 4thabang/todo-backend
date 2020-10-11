package db

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	GoBrr bool   `json:"brr"`
}

var persons = []byte(`[
	{
		Name:  "Thabang",
		Age:   20,
		GoBrr: true,
	},
	{
		Name:  "Divine",
		Age:   2,
		GoBrr: true,
	},
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
