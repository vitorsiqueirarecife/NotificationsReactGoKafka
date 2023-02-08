package app

import (
	"github.com/vitorsiqueirarecife/bff/app/message"
)

type Container struct {
	Message message.App
}

func Register() *Container {

	container := Container{
		Message: message.NewApp(),
	}

	return &container

}
