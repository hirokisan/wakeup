package usecase

import (
	"github.com/hirokisan/go-sample-clean-architecture/domain"
)

type UserRepository interface {
	Store(*domain.User) error
	Users(*domain.Users) error
}
