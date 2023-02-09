package user

import (
	"github.com/vitorsiqueirarecife/bff/model"
)

type Store interface {
	GetByCategory(category *model.Category, channel *model.Channel) error
}

type storeImpl struct{}

func NewApp() Store {
	return nil
}

// users mocked (it should be done with consultation to the users microservice)
func (a *storeImpl) GetByCategory(list []model.User, category *model.Category) []model.User {

	userPerCategory := []model.User{}

	for _, user := range list {
		for _, s := range user.Subscribed {
			if s.Id == category.Id {
				userPerCategory = append(userPerCategory, user)
				break
			}
		}
	}

	return userPerCategory
}

// users mocked (it should be done with consultation to the users microservice)
func (a *storeImpl) GetByChannel(list []model.User, channel *model.Channel) []model.User {

	userPerChannel := []model.User{}

	for _, user := range list {
		for _, c := range user.Channels {
			if c.Id == channel.Id {
				userPerChannel = append(userPerChannel, user)
				break
			}
		}
	}

	return userPerChannel
}
