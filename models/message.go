package models

import (
	"chat/dto"
	"errors"
	"time"
)

type Message struct {
	ID        int64 `gorm:"primaryKey"`
	Message   string
	SenderID  int64
	Sender    User `gorm:"foreignKey:SenderID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewMessage(
	message string,
	sender User,
) (Message, error) {
	if message == "" {
		return Message{}, errors.New("message should not be empty")
	}
	if sender.ID == 0 {
		return Message{}, errors.New("sender should not be empty")
	}
	return Message{
		Message: message,
		Sender:  sender,
	}, nil
}

func (m Message) ToMessageResponse() dto.MessageResponse {
	return dto.MessageResponse{
		ID:        m.ID,
		Message:   m.Message,
		Sender:    m.Sender.ToUserResponse(),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
