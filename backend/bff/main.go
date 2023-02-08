package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vitorsiqueirarecife/bff/api"
)

func main() {

	app := fiber.New()

	api.Register(api.Options{
		AppGroup: app,
	})

	app.Listen(":3000")

}
