package teststore

import (
	"github.com/arsu4ka/go_rest_api/internal/model"
	"github.com/arsu4ka/go_rest_api/internal/store"
)

type UserRepo struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepo) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreation(); err != nil {
		return err
	}

	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}

func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	for _, v := range r.users {
		if v.Email == email {
			return v, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (r *UserRepo) FindByID(id int) (*model.User, error) {
	u, found := r.users[id]
	if !found {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
