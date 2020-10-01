package models

import "time"

// User Model
type User struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	Username       string
	CreatedAt      time.Time
	CompletedTodos int
	ActiveTodos    int
}
