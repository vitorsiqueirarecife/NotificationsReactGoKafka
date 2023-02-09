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
}

type appImpl struct {
	Connection sarama.Consumer
	Store      *store.Container
	Topic      string
}

type Options struct {
	Connection sarama.Consumer
	Store      *store.Container
	Topic      string
}

func NewApp(opt Options) App {
	return &appImpl{
		Topic:      opt.Topic,
		Connection: opt.Connection,
		Store:      opt.Store,
	}
}

func (a *appImpl) Listen() error {
	consumer, err := a.Connection.ConsumePartition(a.Topic, 0, sarama.OffsetOldest)
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
