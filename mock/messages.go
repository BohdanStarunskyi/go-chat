package mock

import (
	"chat/dto"
	"time"
)

type MockMessageService struct{}

func (m *MockMessageService) GetAllMessages(offset int, limit int) ([]dto.MessageResponse, error) {
	return []dto.MessageResponse{
		{
			ID:      1,
			Message: "Test message",
			Sender: dto.UserResponse{
				ID:   1,
				Name: "John",
			},
			CreatedAt: time.Now(),
		},
	}, nil
}

func (m *MockMessageService) HandleIncomingMessage(userID int64, req dto.MessageRequest) (dto.MessageResponse, error) {
	return dto.MessageResponse{
		ID:      2,
		Message: req.Message,
		Sender: dto.UserResponse{
			ID:   userID,
			Name: "MockUser",
		},
		CreatedAt: time.Now(),
	}, nil
}
