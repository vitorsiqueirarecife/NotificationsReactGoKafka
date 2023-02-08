package store

import (
	"github.com/vitorsiqueirarecife/bff/store/user"
)

type Container struct {
	User user.Store
}

func Register() *Container {

	container := Container{
		User: user.NewApp(),
	}

	return &container

}
