package routes

import (
	"chat/app"
	"chat/controllers"
	"chat/middleware"
	"chat/sockets"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, app *app.App) {
	e.GET("/ping", controllers.Healthcheck)
	e.POST("/login", app.AuthController.Login)
	e.POST("/signup", app.AuthController.Signup)
	e.GET("/chat", func(c echo.Context) error {
		sockets.HandleWebSocket(c.Response(), c.Request())
		return nil
	})
	authGroup := e.Group("", middleware.AuthMiddleware)
	authGroup.GET("/messages", app.MessageController.GetAllMessages)

}
