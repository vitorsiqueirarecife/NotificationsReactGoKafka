package api

import (
	"github.com/gofiber/fiber"
	"github.com/vitorsiqueirarecife/bff/api/message"
	"github.com/vitorsiqueirarecife/bff/app"
)

type Options struct {
	Fiber *fiber.App
	App   *app.Container
}

func Register(opts Options) {

	route := opts.Fiber.Group("/api/v1")

	message.NewAPI(&route, opts.App)

}
