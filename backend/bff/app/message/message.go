package message

import (
	"github.com/Shopify/sarama"
	"github.com/vitorsiqueirarecife/bff/model"
)

type App interface {
	SendMessage(message *model.Message) error
}

type appImpl struct {
	Connection sarama.SyncProducer
}

type Options struct {
	Connection sarama.SyncProducer
}

func NewApp(opt Options) App {
	return &appImpl{
		Connection: opt.Connection,
	}
}

func (a *appImpl) SendMessage(message *model.Message) error {

	defer a.Connection.Close()
	// cmtInBytes, err := json.Marshal(message)

	return nil
}
