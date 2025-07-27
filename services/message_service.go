package services

import (
	"chat/database"
	"chat/dto"
	"chat/models"
	"errors"
)

type MessageService interface {
	GetAllMessages(offset int, limit int) ([]dto.MessageResponse, error)
	HandleIncomingMessage(userID int64, req dto.MessageRequest) (dto.MessageResponse, error)
}

type messageService struct{}

func NewMessageService() MessageService {
	return &messageService{}
}

func (m *messageService) GetAllMessages(offset int, limit int) ([]dto.MessageResponse, error) {
	var messages []models.Message
	result := database.DB.Preload("Sender").Offset(offset).Limit(limit).Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}
	var messageDtos = make([]dto.MessageResponse, len(messages))
	for i, value := range messages {
		messageDtos[i] = value.ToMessageResponse()
	}
	return messageDtos, nil
}

func (m *messageService) HandleIncomingMessage(userID int64, req dto.MessageRequest) (dto.MessageResponse, error) {
	switch req.Action {
	case dto.MessageActionAdd:
		return m.addMessage(userID, req.Message)
	case dto.MessageActionEdit:
		return m.editMessage(userID, req.ID, req.Message)
	case dto.MessageActionDelete:
		return dto.MessageResponse{ID: req.ID}, m.deleteMessageByID(userID, req.ID)
	default:
		return dto.MessageResponse{}, errors.New("unknown message action")
	}
}

func (m *messageService) addMessage(userID int64, text string) (dto.MessageResponse, error) {
	user, err := GetUser(userID)
	if err != nil {
		return dto.MessageResponse{}, errors.New("user doesn't exist")
	}

	message, err := models.NewMessage(text, user)
	if err != nil {
		return dto.MessageResponse{}, err
	}

	result := database.DB.Save(&message)
	return message.ToMessageResponse(), result.Error
}

func (m *messageService) editMessage(userID int64, id int64, newContent string) (dto.MessageResponse, error) {
	var message models.Message
	result := database.DB.First(&message, id)
	if result.Error != nil {
		return dto.MessageResponse{}, result.Error
	}
	if message.SenderID != userID {
		return dto.MessageResponse{}, errors.New("user not authorized to edit this message")
	}

	message.Message = newContent
	result = database.DB.Save(&message)
	return message.ToMessageResponse(), result.Error
}

func (m *messageService) deleteMessageByID(userID int64, id int64) error {
	var message models.Message
	result := database.DB.First(&message, id)
	if result.Error != nil {
		return result.Error
	}

	if message.SenderID != userID {
		return errors.New("user not authorized to delete this message")
	}

	result = database.DB.Delete(&models.Message{}, id)
	return result.Error
}
