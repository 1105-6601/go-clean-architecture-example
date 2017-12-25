package main

import (
	"github.com/1105-6601/go-clean-architecture-example/src/infrastructure/config"
	"github.com/1105-6601/go-clean-architecture-example/src/infrastructure/router"
	"github.com/1105-6601/go-clean-architecture-example/src/infrastructure/database"
)

func main() {
	config.NewLoader().Load()
	handler := database.NewMysqlHandler()
	defer handler.Close()
	router.Engine(handler).Run()
}
