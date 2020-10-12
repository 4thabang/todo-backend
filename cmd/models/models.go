package models

// Todos is a structural json map that allows us to organise our SQL table columns.
type Todos struct {
	ID        uint   `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	Owner     string `json:"owner"`
}
