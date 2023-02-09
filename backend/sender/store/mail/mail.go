package mail

import (
	"github.com/vitorsiqueirarecife/bff/model"
)

type Store interface {
	Save(message model.Message) error
}

type storeImpl struct{}

func NewApp() Store {
	return nil
}

func (a *storeImpl) Save(message model.Message) error {
	return nil
}
