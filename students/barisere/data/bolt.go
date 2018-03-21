package data

import (
	_ "github.com/boltdb/bolt"
)

// BoltStorage stores and manipulates Todo items in a Bolt database
type BoltStorage struct {
	name string
}

// Save a Todo item to a Bolt database
func (bs BoltStorage) Save(todo Todo) error {
	return nil
}

// Remove a Todo item from a Bolt database
func (bs BoltStorage) Remove(todo Todo) error {
	return nil
}

// Do marks a Todo item 'done' in a Bolt database
func (bs BoltStorage) Do(todo Todo) error {
	return nil
}

// List returns a Todo list with items that match the
// fields of the Todo items passed in.
//
// The specificity of the matching is, in decreasing order of priority,
// ID > Status > Description
func (bs BoltStorage) List(todo ...Todo) []Todo {
	return nil
}
