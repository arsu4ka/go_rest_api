package sqlstore

import (
	"database/sql"

	"github.com/arsu4ka/go_rest_api/internal/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db       *sql.DB
	userRepo *UserRepo
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() store.UserRepo {
	if s.userRepo != nil {
		return s.userRepo
	}

	s.userRepo = &UserRepo{
		store: s,
	}

	return s.userRepo
}
