package main

import (
	"github.com/Shopify/sarama"
	"github.com/vitorsiqueirarecife/sender/app"
)

func main() {

	topic := "comments"
	worker, err := connectConsumer([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}

	app := app.Register(app.Options{
		Topic:      topic,
		Connection: worker,
	})

	err = app.Message.Listen()
	if err != nil {
		panic(err)
	}

	if err := worker.Close(); err != nil {
		panic(err)
	}

}

func connectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create new consumer
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
