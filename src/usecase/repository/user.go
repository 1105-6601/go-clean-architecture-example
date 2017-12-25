package repository

import "github.com/1105-6601/go-clean-architecture-example/src/domain"

type UserRepository interface {
	One (id domain.UserID) (*domain.User, error)
}
