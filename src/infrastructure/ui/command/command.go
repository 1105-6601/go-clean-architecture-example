package main

import (
	"os"
	"fmt"
	"flag"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"

	"github.com/1105-6601/go-clean-architecture-example/src/infrastructure/config"
	"github.com/1105-6601/go-clean-architecture-example/src/interface/controller"
	"github.com/1105-6601/go-clean-architecture-example/src/infrastructure/database"
	"github.com/1105-6601/go-clean-architecture-example/src/infrastructure/ui/context"
	"github.com/1105-6601/go-clean-architecture-example/src/infrastructure/logger"
	"github.com/1105-6601/go-clean-architecture-example/src/util"
)

var (
	flagSayHello       = flag.Bool("say-hello", false, "Greet.")
	flagMigrate        = flag.Bool("migrate",   false, "Execute database migration.")
	flagMigrationStep  = flag.Int ("step",      1,     "Specify the desired step of migrations.")
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

	if *flagMigrate {

		dir :=  fmt.Sprintf("file://%s/migrations", util.RootDir())
		m, err := migrate.New(dir, database.GetFullDSN())

		if err != nil {
			printer.Error(err.Error())
			return
		}

		err = m.Steps(*flagMigrationStep)
		return
	}
}
