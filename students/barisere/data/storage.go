package data

// Storage provides functionality for persisting TODOs
type Storage interface {
	Save(Todo) error
	Remove(Todo) error
	Do(Todo) error
	List() ([]Todo, error)
	Close() error
}

// store embeds a Storage object
type store struct {
	Storage
}

// NewStorage creates a Storage with the given interface. s should implement the Storage interface
func NewStorage() (Storage, error) {
	storage, err := newBoltStorage()
	if err != nil {
		return nil, err
	}
	return store{storage}, nil
}
