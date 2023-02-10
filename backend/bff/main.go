package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vitorsiqueirarecife/bff/api"
	"github.com/vitorsiqueirarecife/bff/app"
	"github.com/vitorsiqueirarecife/bff/store"
)

func main() {

	store := store.Register()

	app := app.Register(app.Options{
		Store: store,
	})
	fiber := fiber.New()

	api.Register(api.Options{
		Fiber: fiber,
		App:   app,
	})

	fiber.Listen(":3000")

}
