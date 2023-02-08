package main

import (
	"github.com/gofiber/fiber"
	"github.com/vitorsiqueirarecife/bff/api"
	"github.com/vitorsiqueirarecife/bff/app"
)

func main() {

	App := app.Register()
	Fiber := fiber.New()

	api.Register(api.Options{
		Fiber: Fiber,
		App:   App,
	})

	Fiber.Listen(":3000")

}
