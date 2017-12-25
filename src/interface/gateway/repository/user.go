package repository

import (
	"github.com/1105-6601/go-clean-architecture-example/src/interface/gateway/database"
	"github.com/1105-6601/go-clean-architecture-example/src/domain"
)

type UserRepository struct {
	database.Handler
}

func (repository *UserRepository) One(id domain.UserID) (*domain.User, error) {

	user := new(domain.User)

	err := repository.Handler.FindByID(user, int(id))

	return user, err
}
