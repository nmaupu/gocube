package main

import (
	"github.com/nmaupu/gocube/cli"
)

const (
	AppName = "gocube"
	AppDesc = "Rubik's cube utilities written in Go"
)

var (
	AppVersion string
)

func main() {
	if AppVersion == "" {
		AppVersion = "master"
	}

	cli.Process(AppName, AppDesc, AppVersion)
}
