package sqlstore

import (
	"database/sql"
	"github.com/alex-dev-master/golang-rest-example/internal/app/store"
)

// Store ...
type Store struct {
	Db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		Db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		Store: s,
	}

	return s.userRepository
}
