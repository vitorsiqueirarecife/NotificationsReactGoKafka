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
	SendTopic(topic string, message model.Message, user model.User) error
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
			err := a.SendTopic("sms", message, user)
			if err != nil {
				fmt.Println(err)
			}
		}
		wg.Done()
	}()

	go func() {
		users = a.Store.User.GetByChannel(users, mocks.Email.Id)
		for _, user := range users {
			err := a.SendTopic("email", message, user)
			if err != nil {
				fmt.Println(err)
			}
		}
		wg.Done()
	}()

	go func() {
		users = a.Store.User.GetByChannel(users, mocks.PushNotification.Id)
		for _, user := range users {
			err := a.SendTopic("push", message, user)
			if err != nil {
				fmt.Println(err)
			}
		}
		wg.Done()
	}()

	wg.Wait()

	return nil
}

func (a *appImpl) SendTopic(topic string, message model.Message, user model.User) error {
	connection, err := a.ConnectProducer([]string{"localhost:9092"})
	if err != nil {
		return err
	}

	defer connection.Close()

	bytes, err := json.Marshal(model.Message{
		CategoryID: message.CategoryID,
		Text:       message.Text,
		User:       user,
	})
	if err != nil {
		fmt.Println(err)
	}

	_, _, err = connection.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(bytes),
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

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
