package usecase

import (
	"github.com/hirokisan/go-sample-clean-architecture/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (i *UserInteractor) Add(u domain.User) error {
	_, err := i.UserRepository.Store(u)
}

func (i *UserInteractor) Users(domain.Users) (domain.Users, error) {
	_, err := i.UserRepository.FindAll(domain.Users)
}
