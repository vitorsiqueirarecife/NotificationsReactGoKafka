package app

import (
	"github.com/Shopify/sarama"
	"github.com/vitorsiqueirarecife/bff/app/message"
)

type Container struct {
	Message message.App
}

type Options struct {
	Connection sarama.SyncProducer
}

func Register(ops Options) *Container {

	container := Container{
		Message: message.NewApp(message.Options{
			Connection: ops.Connection,
		}),
	}

	return &container

}
