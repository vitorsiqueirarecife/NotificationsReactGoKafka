package message

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
	"github.com/vitorsiqueirarecife/bff/model"
	"github.com/vitorsiqueirarecife/sender/store"
)

type App interface {
	Listen() error
	ConnectConsumer(brokersUrl []string) (sarama.Consumer, error)
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

func (a *appImpl) Listen() error {

	topic := "messages"
	connection, err := a.ConnectConsumer([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	consumer, err := connection.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		return err
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	count := 0

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case received := <-consumer.Messages():
				message := model.Message{}
				json.Unmarshal(received.Value, &message)
				a.Store.Message.Save(message)
				count++
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", count, string(received.Topic), message.Text)
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", count, "messages")
	return nil
}

func (a *appImpl) ConnectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
