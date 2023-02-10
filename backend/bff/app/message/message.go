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
	Send(message model.Message) error
	SendTopic(topic string, bytes []byte) error
	ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error)
}

type appImpl struct {
	Store *store.Container
}

type Options struct {
	Store *store.Container
}

func NewApp(opt Options) App {
	return &appImpl{
		Store: opt.Store,
	}
}

func (a *appImpl) Send(message model.Message) error {

	users := []model.User{}
	users = a.Store.User.GetByCategory(*mocks.Users, message.CategoryID)

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		users = a.Store.User.GetByChannel(users, mocks.Sms.Id)
		for _, user := range users {
			fmt.Println("send-sms", user.Name, user.Subscribed, user.Channels)
			messageBytes, _ := json.Marshal(model.Message{
				CategoryID: message.CategoryID,
				Text:       message.Text,
				Target:     user.PhoneNumber,
			})
			a.SendTopic(mocks.Sms.Name, messageBytes)
		}
		wg.Done()
	}()

	go func() {
		users = a.Store.User.GetByChannel(users, mocks.Email.Id)
		for _, user := range users {
			fmt.Println("send-email: ", user.Name, user.Subscribed, user.Channels)
			messageBytes, _ := json.Marshal(model.Message{
				CategoryID: message.CategoryID,
				Text:       message.Text,
				Target:     user.Email,
			})
			a.SendTopic(mocks.Email.Name, messageBytes)
		}
		wg.Done()
	}()

	go func() {
		users = a.Store.User.GetByChannel(users, mocks.PushNotification.Id)
		for _, user := range users {
			fmt.Println("send-push-notification", user.Name, user.Subscribed, user.Channels)
			messageBytes, _ := json.Marshal(model.Message{
				CategoryID: message.CategoryID,
				Text:       message.Text,
				Target:     user.PhoneNumber,
			})
			a.SendTopic(mocks.PushNotification.Name, messageBytes)
		}
		wg.Done()
	}()

	wg.Wait()

	return nil
}

func (a *appImpl) SendTopic(topic string, bytes []byte) error {

	connection, err := a.ConnectProducer([]string{"localhost:3000"})
	if err != nil {
		return err
	}

	defer connection.Close()

	partition, offset, err := connection.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(bytes),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Sent to topic (%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}

func (a *appImpl) ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
