package user

import (
	"github.com/vitorsiqueirarecife/bff/model"
)

type Store interface {
	GetByCategory(list []model.User, categoryId string) []model.User
	GetByChannel(list []model.User, channelId string) []model.User
}

type storeImpl struct{}

func NewApp() Store {
	return nil
}

// users mocked (it should be done with consultation to the users microservice)
func (a *storeImpl) GetByCategory(list []model.User, categoryId string) []model.User {

	userPerCategory := []model.User{}

	for _, user := range list {
		for _, s := range user.Subscribed {
			if s.Id == categoryId {
				userPerCategory = append(userPerCategory, user)
				break
			}
		}
	}

	return userPerCategory
}

// users mocked (it should be done with consultation to the users microservice)
func (a *storeImpl) GetByChannel(list []model.User, channelId string) []model.User {

	userPerChannel := []model.User{}

	for _, user := range list {
		for _, c := range user.Channels {
			if c.Id == channelId {
				userPerChannel = append(userPerChannel, user)
				break
			}
		}
	}

	return userPerChannel
}
