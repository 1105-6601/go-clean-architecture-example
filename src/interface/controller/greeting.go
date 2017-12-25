package controller

import (
	"net/http"
	"strconv"

	"github.com/1105-6601/go-clean-architecture-example/src/usecase"
	"github.com/1105-6601/go-clean-architecture-example/src/interface/gateway/database"
	"github.com/1105-6601/go-clean-architecture-example/src/interface/gateway/repository"
	"github.com/1105-6601/go-clean-architecture-example/src/interface/controller/context"
	"github.com/1105-6601/go-clean-architecture-example/src/usecase/logger"
	"github.com/1105-6601/go-clean-architecture-example/src/domain"
	"github.com/1105-6601/go-clean-architecture-example/src/util"
)

type GreetingController struct {
	Interactor usecase.GreetingInteractor
}

func NewGreetingController(handler database.Handler, printer logger.Printer) *GreetingController {
	return &GreetingController{
		Interactor: usecase.GreetingInteractor{
			UserRepository: &repository.UserRepository{
				Handler: handler,
			},
			Printer:  printer,
		},
	}
}

func (controller *GreetingController) RestSayHello(c context.HTTP) {

	userID, err := strconv.Atoi(c.Param("userID"))

	if err != nil {
		c.JSON(http.StatusBadRequest, util.H{"message": "Bad request."})
		return
	}

	phrase, _ := controller.Interactor.MakePhrase(domain.UserID(userID))

	c.JSON(http.StatusOK, util.H{"message": phrase})
}


func (controller *GreetingController) CmdSayHello(c context.Command) {

	UserID := domain.UserID(0)

	switch  len(c.Args()) {
	case 1:
		UserID = domain.UserID(c.Args()[0].(int))
	}

	phrase, _ := controller.Interactor.MakePhrase(UserID)

	c.Echo(phrase)
}
