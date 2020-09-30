package models

// Todo Model
type Todo struct {
	ID         uint `gorm:"primaryKey"`
	Todo       string
	IsComplete bool
}
