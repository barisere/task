package data

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

var (
	pendingBucket = []byte("todos")
	doneBucket    = []byte("done")
	dbFile        = "bolt.db"
)

// boltStorage stores and manipulates Todo items in a Bolt database
type boltStorage struct {
	name string
	db   *bolt.DB
}

// newBoltStorage constructs a boltStorage with a db connection opened
func newBoltStorage() (*boltStorage, error) {
	db, err := bolt.Open(dbFile, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &boltStorage{
		name: dbFile,
		db:   db,
	}, nil
}

// Save a Todo item to a Bolt database
func (bs boltStorage) Save(todo Todo) error {
	err := bs.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(pendingBucket)
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
		key, value := itob(next), yamlBytes

		if err := bucket.Put(key, value); err != nil {
			return err
		}
		return nil
	})
	return err
}

// Remove a Todo item from a Bolt database
func (bs boltStorage) Remove(todo Todo) error {
	key := itob(todo.ID)
	err := bs.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(pendingBucket)
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
func (bs boltStorage) Do(todo Todo) error {
	key := itob(todo.ID)
	err := bs.db.Update(func(tx *bolt.Tx) error {
		pending, err := tx.CreateBucketIfNotExists(pendingBucket)
		if err != nil {
			return err
		}
		done, err := tx.CreateBucketIfNotExists(doneBucket)
		if err != nil {
			return err
		}
		if fetched := pending.Get(key); fetched != nil {
			oldTodo, err := fromYAMLBytes(fetched)
			if err != nil {
				return err
			}
			newTodoBytes, err := toYAMLBytes(oldTodo.Do())
			if err != nil {
				return err
			}
			if err = pending.Delete(key); err != nil {
				return err
			}
			if err = done.Put(key, newTodoBytes); err != nil {
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
func (bs boltStorage) List() ([]Todo, error) {
	var todos []Todo
	err := bs.db.View(func(tx *bolt.Tx) error {
		pendings := tx.Bucket(pendingBucket)
		err := pendings.ForEach(func(k, v []byte) error {
			todoItem, err := fromYAMLBytes(v)
			if err != nil {
				return err
			}
			todos = append(todos, todoItem)
			return nil
		})
		dones := tx.Bucket(doneBucket)
		err = dones.ForEach(func(k, v []byte) error {
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

func (bs boltStorage) Close() error {
	return bs.db.Close()
}

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
