package app

import (
	"chat/controllers"
	"chat/services"
)

type App struct {
	AuthService       *services.AuthService
	AuthController    *controllers.AuthController
	MessageService    *services.MessageService
	MessageController *controllers.MessagesController
}

func NewApp() *App {
	authService := services.NewAuthService()
	authController := controllers.NewAuthController(authService)

	messageService := services.NewMessageService()
	messageController := controllers.NewMessageController(messageService)
	return &App{
		AuthService:       &authService,
		AuthController:    authController,
		MessageService:    &messageService,
		MessageController: messageController,
	}
}
