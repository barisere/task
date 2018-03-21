package data

// Status of a Todo
const (
	Pending = iota
	Done
	Removed
)

// Todo is a task to be done sometime
type Todo struct {
	ID          int
	Status      int
	Description string
}

// Storage provides functionality for persisting TODOs
type Storage interface {
	Save(Todo) error
	Remove(Todo) error
	Do(Todo) error
	List() []Todo
}
