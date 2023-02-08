package teststore

import (
	"github.com/arsu4ka/go_rest_api/internal/model"
	"github.com/arsu4ka/go_rest_api/internal/store"
	_ "github.com/lib/pq"
)

type Store struct {
	userRepo *UserRepo
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepo {
	if s.userRepo != nil {
		return s.userRepo
	}

	s.userRepo = &UserRepo{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepo
}
