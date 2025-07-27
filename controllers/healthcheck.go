package controllers

import (
	"chat/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Healthcheck(c echo.Context) error {
	return c.JSON(http.StatusOK, dto.NewMessageResponse("pong!"))
}
