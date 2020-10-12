package db

import (
	"fmt"
	"net/http"
)

// Todos is a structural json map that allows us to organise our SQL table columns.
type Todos struct {
	ID        uint   `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	Owner     string `json:"owner"`
}

// CreateData allows us to index our database and add fresh new values.
func CreateData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create Data")
}

// ReadData allows us to query our database and read it's values.
func ReadData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Read Data")
}

// UpdateData allows us to update query entries in our SQL database.
func UpdateData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Update Data")
}

// DeleteData allows us to delete queries in our SQL database.
func DeleteData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete Data")
}
