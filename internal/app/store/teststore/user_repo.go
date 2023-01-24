package teststore

import (
	"github.com/arsu4ka/go_rest_api/internal/app/model"
	"github.com/arsu4ka/go_rest_api/internal/app/store"
)

type UserRepo struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepo) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreation(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
