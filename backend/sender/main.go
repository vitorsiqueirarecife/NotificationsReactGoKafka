package main

import (
	"github.com/vitorsiqueirarecife/sender/app"
	"github.com/vitorsiqueirarecife/sender/store"
)

func main() {

	store := store.Register()

	app := app.Register(app.Options{
		Store: store,
	})

	err := app.Message.Listen()
	if err != nil {
		panic(err)
	}

}
