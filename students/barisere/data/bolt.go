package data

import (
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var bucketName = []byte("todos")

// BoltStorage stores and manipulates Todo items in a Bolt database
type BoltStorage struct {
	Name string
	DB   *bolt.DB
}

// NewBoltStorage constructs a boltStorage with a DB connection opened
func NewBoltStorage(dbFile string) (*BoltStorage, error) {
	if dbFile == "" {
		return nil, errors.New("no database file passed")
	}
	db, err := bolt.Open(dbFile, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &BoltStorage{
		Name: dbFile,
		DB:   db,
	}, nil
}

// Save a Todo item to a Bolt database
func (bs BoltStorage) Save(todo Todo) error {
	err := bs.DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}
		next, err := bucket.NextSequence()
		if err != nil {
			return errors.Wrap(err, "Unable to obtain ID for Todo")
		}
		todo = NewTodo(next, todo.Status, todo.Description)
		yamlBytes, err := toYAMLBytes(todo)
		if err != nil {
			return err
		}
		key, value := []byte(todo.Description), yamlBytes

		if err := bucket.Put(key, value); err != nil {
			return err
		}
		return nil
	})
	return err
}

// Remove a Todo item from a Bolt database
func (bs BoltStorage) Remove(todo Todo) error {
	key := []byte(todo.Description)
	err := bs.DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}
		if err := bucket.Delete(key); err != nil {
			return err
		}
		return nil
	})
	return err
}

// Do marks a Todo item 'done' in a Bolt database
func (bs BoltStorage) Do(todo Todo) error {
	key := []byte(todo.Description)
	err := bs.DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}
		if fetched := bucket.Get(key); fetched != nil {
			oldTodo, err := fromYAMLBytes(fetched)
			if err != nil {
				return err
			}
			newTodoBytes, err := toYAMLBytes(oldTodo.Do())
			if err != nil {
				return err
			}
			if err = bucket.Put(key, newTodoBytes); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// List returns a Todo list with items that match the
// fields of the Todo items passed in.
//
// The specificity of the matching is, in decreasing order of priority,
// ID > Status > Description
func (bs BoltStorage) List(todo ...Todo) ([]Todo, error) {
	var todos []Todo
	err := bs.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		err := bucket.ForEach(func(k, v []byte) error {
			todoItem, err := fromYAMLBytes(v)
			if err != nil {
				return err
			}
			todos = append(todos, todoItem)
			return nil
		})
		return err
	})
	return todos, err
}

func toYAMLBytes(todo Todo) ([]byte, error) {
	return yaml.Marshal(todo)
}

func fromYAMLBytes(yamlBytes []byte) (Todo, error) {
	todo := Todo{}
	err := yaml.Unmarshal(yamlBytes, &todo)
	return todo, err
}
