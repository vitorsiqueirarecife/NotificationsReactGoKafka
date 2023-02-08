package model

type Category struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
