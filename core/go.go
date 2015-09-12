package core

import (
	"runtime"
)

func GoRoot() string {
	return runtime.GOROOT()
}

func GoVersion() string {
	return runtime.Version()
}

