package teststore

import (
	"github.com/dmitryzhvinklis/carfix_dev/internal/app/model"
	"github.com/dmitryzhvinklis/carfix_dev/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return nil
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
