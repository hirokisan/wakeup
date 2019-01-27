package usecase

import (
	"github.com/hirokisan/go-sample-clean-architecture/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (i *UserInteractor) Add(u *domain.User) error {
	err := i.UserRepository.Store(u)
	if err != nil {
		return err
	}
	return nil
}

func (i *UserInteractor) Users(users *domain.Users) error {
	err := i.UserRepository.Users(users)
	if err != nil {
		return err
	}
	return nil
}
