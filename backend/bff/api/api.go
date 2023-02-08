package api

import "github.com/gofiber/fiber/v2"

type Options struct {
	AppGroup *fiber.App
}

func Register(opts Options) {

	api := opts.AppGroup.Group("/api/v1")

	api.Post("/comments", nil)

}
