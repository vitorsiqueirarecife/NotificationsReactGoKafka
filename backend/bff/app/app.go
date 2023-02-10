package app

import (
	"github.com/vitorsiqueirarecife/bff/app/message"
	"github.com/vitorsiqueirarecife/bff/store"
)

type Container struct {
	Message message.App
}

type Options struct {
	Store *store.Container
}

func Register(ops Options) *Container {

	mApp := message.NewApp(message.Options{
		Store: ops.Store,
	})

	container := Container{
		Message: mApp,
	}

	return &container

}
