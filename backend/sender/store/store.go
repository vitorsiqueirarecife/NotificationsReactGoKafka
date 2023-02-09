package store

import "github.com/vitorsiqueirarecife/sender/store/message"

type Container struct {
	Message message.Store
}

func Register() *Container {

	container := Container{
		Message: message.NewApp(),
	}

	return &container

}
