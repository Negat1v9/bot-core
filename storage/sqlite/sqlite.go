package sqlite

import (
	"database/sql"

	"github.com/Negat1v9/bot-core/storage"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db             *sql.DB
	userRepository *UserRepository
}

// Constructor for Storage.
func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

// Method for use repository from storage.
func (s *Storage) User() storage.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		storage: s,
	}
	return s.userRepository
}
