package model

type Message struct {
	CategoryID string `json:"category_id" validate:"required"`
	Text       string `json:"text" validate:"required"`
}
