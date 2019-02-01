package database

import (
	"github.com/hirokisan/go-sample-clean-architecture/domain"
)

type UserRepository struct {
	MongoCollectionHandler
}

func (repo *UserRepository) Store(user *domain.User) error {
	if err := repo.Insert(user); err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Users(users *domain.Users) error {
	if err := repo.FindAll(users); err != nil {
		return err
	}
	return nil
}
