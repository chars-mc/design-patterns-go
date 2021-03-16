package storage

import "fmt"

// Storage contains the DB engine
type Storage struct {
	engine string
}

// NewStorage creates a Storage
func NewStorage(e string) Storage {
	return Storage{e}
}

// Save saves the comment to DB
func (s *Storage) Save(comment string) {
	fmt.Printf("comments saved on DB:%s\n", s.engine)
}
