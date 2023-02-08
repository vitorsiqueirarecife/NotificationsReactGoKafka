package model

type User struct {
	Id          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Subscribed  []Category
	Channels    []Channel
}
