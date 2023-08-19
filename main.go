package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
)

func main() {
	app := pocketbase.New()

	jsvm.MustRegister(app, jsvm.Config{
		HooksWatch: true, // make this false for production
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
