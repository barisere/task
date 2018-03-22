package data

import (
	"path/filepath"
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

func TestNewBoltStorage(t *testing.T) {
	// NewBoltStorage should return a BoltStorage and no errors
	boltdb, err := NewBoltStorage(path)
	if err != nil {
		t.Errorf("Failed to create DB:\n%+v", err)
	}
	defer boltdb.DB.Close()

	// If successful, the pathname of the db should be the same as that passed to NewBoltStorage
	dbpath := filepath.Base(boltdb.DB.Path())
	if dbpath != path {
		t.Errorf("Created unexpected DB: wanted %s, got %s\n", path, dbpath)
	}
}

// func TestSave(t *testing.T) {
// 	boltdb, err := NewBoltStorage(path)
// 	if err != nil {
// 		t.Errorf("Failed to create DB:\n%+v", err)
// 	}
// 	defer boltdb.db.Close()

// 	// if saving to database fails, we should not find 'todo' in database;
// 	// otherwise, todo should be in database
// 	err = boltdb.Save(todo)
// 	if err != nil {
// 		t.Run("a failed save should n")
// 	}

// }
