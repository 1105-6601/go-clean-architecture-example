package router

import (
	"github.com/gin-gonic/gin"

	"github.com/1105-6601/go-clean-architecture-example/src/interface/controller"
	"github.com/1105-6601/go-clean-architecture-example/src/interface/gateway/database"
	"github.com/1105-6601/go-clean-architecture-example/src/infrastructure/ui/context"
	"github.com/1105-6601/go-clean-architecture-example/src/infrastructure/logger"
)

func Engine(handler database.Handler) *gin.Engine {

	engine  := gin.Default()
	printer := logger.NewColorPrinter()

	greetingController := controller.NewGreetingController(handler, printer)

	v1 := engine.Group("/v1")
	{
		v1.GET("/hello/:userID", func(c *gin.Context) { greetingController.RestSayHello(context.NewHTTP(c)) })
	}

	return engine
}
