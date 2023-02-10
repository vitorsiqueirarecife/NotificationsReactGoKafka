package app

import (
	"github.com/vitorsiqueirarecife/sender/app/message"
	"github.com/vitorsiqueirarecife/sender/store"
)

type Container struct {
	Message message.App
}

type Options struct {
	Store *store.Container
}

func Register(ops Options) *Container {

	container := Container{
		Message: message.NewApp(message.Options{
			Store: ops.Store,
		}),
	}

	return &container

}
