package usecase

import (
	"github.com/hirokisan/go-sample-clean-architecture/domain"
)

type UserRepository interface {
	Store(domain.User) error
	FindAll(domain.Users) (domain.Users, error)
}
