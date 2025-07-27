package dto

import (
	"time"
)

type MessageAction string

const (
	MessageActionAdd    MessageAction = "add"
	MessageActionEdit   MessageAction = "edit"
	MessageActionDelete MessageAction = "delete"
)

type MessageRequest struct {
	ID      int64         `json:"id" validate:"-"`
	Message string        `json:"message" validate:"required"`
	Action  MessageAction `json:"action" validate:"required"`
}

type MessageResponse struct {
	ID        int64        `json:"id"`
	Message   string       `json:"message"`
	Sender    UserResponse `json:"sender"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
}
