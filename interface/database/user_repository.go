package database

import (
	"github.com/hirokisan/clean/domain"
)

type UserRepository struct {
	MongoHandler
}

func (repo *UserRepository) Store(user domain.User) error {
	if err := repo.Insert(user); err != nil {
		return err
	}
	return nil
}
