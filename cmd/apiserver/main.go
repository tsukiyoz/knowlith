package main

import (
	"github.com/tsukiyoz/knowlith/cmd/apiserver/app"
	"os"
)

func main() {
	command := app.NewAPIServerCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
