package api

import (
	"github.com/gofiber/fiber"
	"github.com/vitorsiqueirarecife/bff/api/message"
	"github.com/vitorsiqueirarecife/bff/app"
)

type Options struct {
	AppGroup *fiber.App
	app      *app.Container
}

func Register(opts Options) {

	route := opts.AppGroup.Group("/api/v1")

	message.NewAPI(&route, opts.app)

}
