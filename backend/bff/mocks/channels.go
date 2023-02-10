package mocks

import "github.com/vitorsiqueirarecife/bff/model"

var Sms = model.Channel{
	Id:   "1",
	Name: "messages-sms",
}

var Email = model.Channel{
	Id:   "2",
	Name: "messages-email",
}

var PushNotification = model.Channel{
	Id:   "3",
	Name: "messages-push-notification",
}
