package message

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/vitorsiqueirarecife/bff/model"
	"github.com/vitorsiqueirarecife/bff/store"
)

type App interface {
	SendMessage(message *model.Message) error
}

type appImpl struct {
	Connection sarama.SyncProducer
	Store      *store.Container
}

type Options struct {
	Connection sarama.SyncProducer
	Store      *store.Container
}

func NewApp(opt Options) App {
	return &appImpl{
		Connection: opt.Connection,
		Store:      opt.Store,
	}
}

func (a *appImpl) SendMessage(message *model.Message) error {
	messageBytes, err := json.Marshal(message)

	defer a.Connection.Close()

	topic := "topic_name"
	msg := &sarama.ProducerMessage{
		Topic: "topic_name",
		Value: sarama.StringEncoder(messageBytes),
	}

	partition, offset, err := a.Connection.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}
