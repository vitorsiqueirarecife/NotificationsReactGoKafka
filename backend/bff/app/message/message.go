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
	SendTopic(topic string, message model.Message, users []model.User) error
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
		usersSms := a.Store.User.GetByChannel(users, mocks.Sms.Id)
		err := a.SendTopic("sms", message, usersSms)
		if err != nil {
			fmt.Println(err)
		}
		wg.Done()
	}()

	go func() {
		usersEmail := a.Store.User.GetByChannel(users, mocks.Email.Id)
		err := a.SendTopic("email", message, usersEmail)
		if err != nil {
			fmt.Println(err)
		}
		wg.Done()
	}()

	go func() {
		usersPush := a.Store.User.GetByChannel(users, mocks.PushNotification.Id)
		err := a.SendTopic("push", message, usersPush)
		if err != nil {
			fmt.Println(err)
		}
		wg.Done()
	}()

	wg.Wait()

	return nil
}

func (a *appImpl) SendTopic(topic string, message model.Message, users []model.User) error {
	connection, err := a.ConnectProducer([]string{"localhost:9092"})
	if err != nil {
		return err
	}

	defer connection.Close()

	topicMessages := []*sarama.ProducerMessage{}

	for _, user := range users {

		bytes, err := json.Marshal(model.Message{
			CategoryID: message.CategoryID,
			Text:       message.Text,
			User:       user,
		})

		if err != nil {
			fmt.Println(err)
		}

		topicMessages = append(topicMessages, &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(bytes),
		})

	}

	err = connection.SendMessages(topicMessages)

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
