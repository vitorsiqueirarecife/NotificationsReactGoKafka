package message

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/vitorsiqueirarecife/bff/mocks"
	"github.com/vitorsiqueirarecife/bff/model"
	"github.com/vitorsiqueirarecife/bff/store"
)

type App interface {
	Send(message *model.Message) error
	SendTopic(topic string, bytes []byte) error
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

func (a *appImpl) Send(message *model.Message) error {

	defer a.Connection.Close()

	messageBytes, _ := json.Marshal(message.Text)
	users := []model.User{}
	users = a.Store.User.GetByCategory(*mocks.Users, message.CategoryID)

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		users = a.Store.User.GetByChannel(users, mocks.Sms.Id)
		if len(users) > 0 {
			a.SendTopic(mocks.Sms.Name, messageBytes)
		}
		wg.Done()
	}()

	go func() {
		users = a.Store.User.GetByChannel(users, mocks.Email.Id)
		if len(users) > 0 {
			a.SendTopic(mocks.Email.Name, messageBytes)
		}
		wg.Done()
	}()

	go func() {
		users = a.Store.User.GetByChannel(users, mocks.PushNotification.Id)
		if len(users) > 0 {
			a.SendTopic(mocks.PushNotification.Name, messageBytes)
		}
		wg.Done()
	}()

	wg.Wait()

	return nil
}

func (a *appImpl) SendTopic(topic string, bytes []byte) error {

	partition, offset, err := a.Connection.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(bytes),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Sent to topic (%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}
