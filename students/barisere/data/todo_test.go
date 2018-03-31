package data

import (
	"testing"
)

var todo = Todo{
	ID:          1,
	Status:      Pending,
	Description: "TestTodo",
}
var path = "bolt.db"

func TestFromYAMLBytes(t *testing.T) {

	yamlBytes, err := toYAMLBytes(todo)
	if err != nil {
		t.Errorf("%+v", err)
	}
	todoFromYAML, err := fromYAMLBytes(yamlBytes)
	if err != nil {
		t.Errorf("Converting from YAML failed with error:\n%+v", err)
	}
	if todo != todoFromYAML {
		t.Errorf("Unexpected Todo gotten:\nwant %+v\ngot %+v", todo, todoFromYAML)
	}
}

func TestDo(t *testing.T) {
	doneTodo := todo.Do()
	if doneTodo.Status != Done {
		t.Errorf("Expected todo.Status to be Done, got %v\n", doneTodo.Status)
	}
}
