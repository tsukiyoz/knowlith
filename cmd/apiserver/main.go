package main

import (
	"os"

	"github.com/tsukiyoz/knowlith/cmd/apiserver/app"
)

func main() {
	command := app.NewAPIServerCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
