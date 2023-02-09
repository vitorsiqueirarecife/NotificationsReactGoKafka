package mocks

import "github.com/vitorsiqueirarecife/bff/model"

var Sms = model.Channel{
	Id:   "1",
	Name: "SMS",
}

var Email = model.Channel{
	Id:   "2",
	Name: "E-Mail",
}

var PushNotification = model.Channel{
	Id:   "3",
	Name: "Push Notification",
}
