package message

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/vitorsiqueirarecife/bff/app"
	"github.com/vitorsiqueirarecife/bff/model"
)

type apiImpl struct {
	apps *app.Container
	fi   fiber.Router
}

func NewAPI(f *fiber.Router, apps *app.Container) {
	api := apiImpl{
		apps: apps,
		fi:   *f,
	}

	api.fi.Post("/send-message", api.SendMessage)
}

func (a *apiImpl) SendMessage(c *fiber.Ctx) {

	message := new(model.Message)

	if err := c.BodyParser(message); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	a.apps.Message.SendMessage(message)
}
