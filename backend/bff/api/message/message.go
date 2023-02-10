package message

import (
	"log"

	"github.com/gofiber/fiber/v2"
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

func (a *apiImpl) SendMessage(c *fiber.Ctx) error {

	message := new(model.Message)

	if err := c.BodyParser(message); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	err := a.apps.Message.Send(message)

	err = c.JSON(&fiber.Map{
		"success":  true,
		"message":  "Message send successfully",
		"messages": message,
	})
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
	}

	return nil
}
