package utility

import (
	"path/filepath"
	"runtime"
)

func CallerDir() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "."
	}
	return filepath.Dir(file)
}
