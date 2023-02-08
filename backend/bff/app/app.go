package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vitorsiqueirarecife/bff/app/message"
)

type Container struct {
	Message message.App
}

type Options struct {
	AppGroup *fiber.App
}

func Register(opts Options) *Container {

	container := Container{
		Message: message.NewApp(),
	}

	return &container

}
