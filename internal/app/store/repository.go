package store

import "github.com/dmitryzhvinklis/carfix_dev/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
