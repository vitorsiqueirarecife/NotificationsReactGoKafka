package message

import (
	"log"
	"os"

	"github.com/vitorsiqueirarecife/bff/model"
)

type Store interface {
	Save(message model.Message) error
}

type storeImpl struct{}

func NewApp() Store {
	return &storeImpl{}
}

func (a *storeImpl) Save(message model.Message) error {

	f, err := os.OpenFile("MY-LOGS-HERE", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	log.SetOutput(f)
	log.Println(message.CategoryID, message.Target, message.Text)

	return nil
}
