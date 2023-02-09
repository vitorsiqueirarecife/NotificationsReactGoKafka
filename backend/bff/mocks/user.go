package mocks

import "github.com/vitorsiqueirarecife/bff/model"

var jose = model.User{
	Id:          "1",
	Name:        "José",
	Email:       "jose@gmail.com",
	PhoneNumber: "558199813918",
	Subscribed:  []model.Category{Sports, Movies},
	Channels:    []model.Channel{Sms, Email},
}

var caio = model.User{
	Id:          "2",
	Name:        "Caio",
	Email:       "caio@gmail.com",
	PhoneNumber: "558199813919",
	Subscribed:  []model.Category{Sports, Finance},
	Channels:    []model.Channel{Sms, PushNotification},
}

var lucas = model.User{
	Id:          "3",
	Name:        "Lucas",
	Email:       "lucas@gmail.com",
	PhoneNumber: "558199813916",
	Subscribed:  []model.Category{Finance},
	Channels:    []model.Channel{Email, Sms},
}

var diego = model.User{
	Id:          "4",
	Name:        "diego",
	Email:       "diego@gmail.com",
	PhoneNumber: "558199813911",
	Subscribed:  []model.Category{Finance},
	Channels:    []model.Channel{Email},
}

var mario = model.User{
	Id:          "5",
	Name:        "mario",
	Email:       "mario@gmail.com",
	PhoneNumber: "558199813912",
	Subscribed:  []model.Category{Movies, Finance, Sports},
	Channels:    []model.Channel{PushNotification},
}

var joao = model.User{
	Id:          "6",
	Name:        "joao",
	Email:       "joao@gmail.com",
	PhoneNumber: "558199813913",
	Subscribed:  []model.Category{Movies, Sports},
	Channels:    []model.Channel{PushNotification, Sms},
}

var Users = &[]model.User{jose, lucas, caio, diego, mario, joao}
