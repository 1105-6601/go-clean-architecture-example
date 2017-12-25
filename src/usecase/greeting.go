package usecase

import (
	"fmt"

	"github.com/1105-6601/go-clean-architecture-example/src/usecase/repository"
	"github.com/1105-6601/go-clean-architecture-example/src/usecase/logger"
	"github.com/1105-6601/go-clean-architecture-example/src/domain"
)

type GreetingInteractor struct {
	UserRepository repository.UserRepository
	Printer        logger.Printer
}

func (interactor *GreetingInteractor) MakePhrase(userID domain.UserID) (string, error) {

	user, err := interactor.UserRepository.One(userID)

	phrase := ""

	if err == nil {
		phrase = fmt.Sprintf("Hello %s!", user.FullName())
	} else {
		interactor.Printer.Error("Failed to find user.")
	}

	return phrase, err
}
