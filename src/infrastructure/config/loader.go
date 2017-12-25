package config

import (
	"github.com/joho/godotenv"
	"path/filepath"
	"runtime"
	"log"
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

	_, file, _, _ := runtime.Caller(0)

	envFile := ".env"

	if loader.Env != "" {
		envFile += "." + loader.Env
	}
	dir := filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(file)))) + "/" + envFile

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
