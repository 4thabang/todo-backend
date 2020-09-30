package models

// User Model
type User struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	Username       string
	CompletedTodos int
	ActiveTodos    int
}
