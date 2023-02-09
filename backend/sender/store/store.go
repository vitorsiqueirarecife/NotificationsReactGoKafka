package store

import "github.com/vitorsiqueirarecife/sender/store/mail"

type Container struct {
	Mail mail.Store
}

func Register() *Container {

	container := Container{
		Mail: mail.NewApp(),
	}

	return &container

}
