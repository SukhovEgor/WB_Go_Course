package main

import (
	"log"

	"DistributedGrep/internal/di"
)

func main() {

	app, err := di.Build()

	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}