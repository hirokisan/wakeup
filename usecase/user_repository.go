package usecase

import (
	"github.com/hirokisan/clean/domain"
)

type UserRepository interface {
	Store(domain.User) error
	FindAll(domain.Users) (domain.Users, error)
}
