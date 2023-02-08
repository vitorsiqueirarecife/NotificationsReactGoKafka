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

// users mocked
func (a *storeImpl) GetByCategory(category *model.Category, channel *model.Channel) *[]model.User {

	if category.Id == "1" {
		return &[]model.User{
			model.User{
				Id:          "1",
				Name:        "José",
				Email:       "jose@gmail.com",
				PhoneNumber: "558199813918",
				Subscribed:  []model.Category{},
				Channels:    []model.Channel{},
			},
		}
	}

	if category.Id == "2" {
		return &[]model.User{
			model.User{
				Id:          "2",
				Name:        "Caio",
				Email:       "caio@gmail.com",
				PhoneNumber: "558199813919",
				Subscribed:  []model.Category{},
				Channels:    []model.Channel{},
			},
		}
	}

	if category.Id == "3" {
		return &[]model.User{
			model.User{
				Id:          "3",
				Name:        "Lucas",
				Email:       "lucas@gmail.com",
				PhoneNumber: "558199813916",
				Subscribed:  []model.Category{},
				Channels:    []model.Channel{},
			},
		}
	}

	return nil

}
