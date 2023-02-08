package message

import (
	"encoding/json"

	"github.com/vitorsiqueirarecife/bff/model"
)

type (
	App interface {
		SendMessage(message *model.Message) error
	}
	appImpl struct {
	}
)

func NewApp() App {
	return nil
}

func (a *appImpl) SendMessage(message *model.Message) error {

	cmtInBytes, err := json.Marshal(message)

	return nil
}
