package app

import (
	"github.com/Shopify/sarama"
	"github.com/vitorsiqueirarecife/sender/app/message"
	"github.com/vitorsiqueirarecife/sender/store"
)

type Container struct {
	Message message.App
}

type Options struct {
	Topic      string
	Connection sarama.Consumer
	Store      *store.Container
}

func Register(ops Options) *Container {

	container := Container{
		Message: message.NewApp(message.Options{
			Connection: ops.Connection,
			Store:      ops.Store,
		}),
	}

	return &container

}
