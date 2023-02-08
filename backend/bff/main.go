package main

import (
	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber"
	"github.com/vitorsiqueirarecife/bff/api"
	"github.com/vitorsiqueirarecife/bff/app"
)

func main() {

	brokersUrl := []string{"localhost:9092"}
	connection, _ := connectProducer(brokersUrl)

	app := app.Register(app.Options{
		Connection: &connection,
	})
	fiber := fiber.New()

	api.Register(api.Options{
		Fiber: fiber,
		App:   app,
	})

	fiber.Listen(":3000")

}

func connectProducer(brokersUrl []string) (sarama.SyncProducer, error) {

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
