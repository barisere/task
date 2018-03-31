package data

import yaml "gopkg.in/yaml.v2"

// Status of a Todo
const (
	Pending TodoStatus = iota
	Done
	Removed
)

// TodoStatus indicates the state of a Todo item
type TodoStatus int

// Todo is a task to be done sometime
type Todo struct {
	ID          uint64
	Status      TodoStatus
	Description string
}

// Do marks a Todo item as done
func (t Todo) Do() Todo {
	return NewTodo(t.ID, Done, t.Description)
}

// NewTodo constructs a new Todo item
func NewTodo(ID uint64, status TodoStatus, description string) Todo {
	if status == 0 {
		status = Pending
	}
	return Todo{
		ID:          ID,
		Status:      status,
		Description: description,
	}
}

func toYAMLBytes(todo Todo) ([]byte, error) {
	return yaml.Marshal(todo)
}

func fromYAMLBytes(yamlBytes []byte) (Todo, error) {
	todo := Todo{}
	err := yaml.Unmarshal(yamlBytes, &todo)
	return todo, err
}
