package util

import (
	"path/filepath"
	"runtime"
)

func RootDir() string {

	_, file, _, _ := runtime.Caller(0)
	path := filepath.Dir(filepath.Dir(filepath.Dir(file)))

	return path
}
