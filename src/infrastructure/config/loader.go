package config

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/1105-6601/go-clean-architecture-example/src/util"
)

type Loader struct {
	Env string
}

func NewLoader(env ...interface{}) *Loader {

	loader := new(Loader)

	if len(env) == 0 {
		return loader
	}

	if str, ok := env[0].(string); ok {
		loader.Env = str
	}

	return loader
}

func (loader *Loader) Load() {

	envFile := ".env"

	if loader.Env != "" {
		envFile += "." + loader.Env
	}

	dir := util.RootDir() + "/" + envFile
	err := godotenv.Load(dir)

	if err == nil {
		return
	}

	dir = "/" + envFile
	err = godotenv.Load(dir)

	if err != nil {
		log.Fatal(err)
	}
}
