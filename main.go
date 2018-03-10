package main

import (
	"log"

	"github.com/thenomemac/goshow/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
