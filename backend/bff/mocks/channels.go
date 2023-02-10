package mocks

import "github.com/vitorsiqueirarecife/bff/model"

var Sms = model.Channel{
	Id:   "1",
	Name: "sms",
}

var Email = model.Channel{
	Id:   "2",
	Name: "email",
}

var PushNotification = model.Channel{
	Id:   "3",
	Name: "push-notification",
}
