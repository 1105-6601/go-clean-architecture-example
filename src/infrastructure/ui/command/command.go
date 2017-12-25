package main

import (
	"os"
	"flag"

	"github.com/1105-6601/go-clean-architecture-example/src/infrastructure/config"
	"github.com/1105-6601/go-clean-architecture-example/src/interface/controller"
	"github.com/1105-6601/go-clean-architecture-example/src/infrastructure/database"
	"github.com/1105-6601/go-clean-architecture-example/src/infrastructure/ui/context"
	"github.com/1105-6601/go-clean-architecture-example/src/infrastructure/logger"
)

var (
	flagSayHello = flag.Bool("say-hello", false, "Greet.")
)

func main() {

	config.NewLoader(os.Getenv("APP_ENV")).Load()

	flag.Parse()

	handler := database.NewMysqlHandler()
	defer handler.Close()

	printer := logger.NewColorPrinter()

	if *flagSayHello {

		greetingController := controller.NewGreetingController(handler, printer)
		greetingController.CmdSayHello(context.NewCommand(1))

		return
	}
}
