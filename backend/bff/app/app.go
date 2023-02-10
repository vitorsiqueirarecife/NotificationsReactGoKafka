package app

import (
	"github.com/Shopify/sarama"
	"github.com/vitorsiqueirarecife/bff/app/message"
	"github.com/vitorsiqueirarecife/bff/store"
)

type Container struct {
	Message message.App
}

type Options struct {
	Connection sarama.SyncProducer
	Store      *store.Container
}

func Register(ops Options) *Container {

	mApp := message.NewApp(message.Options{
		Connection: ops.Connection,
		Store:      ops.Store,
	})

	container := Container{
		Message: mApp,
	}

	return &container

}
