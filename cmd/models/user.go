package models

// User Model
type User struct {
	ID             uint
	Name           string
	Username       string
	CompletedTodos int
	ActiveTodos    int
}
