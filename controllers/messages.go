package controllers

import (
	"chat/dto"
	"chat/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MessagesController struct {
	messageService services.MessageService
}

func NewMessageController(messageService services.MessageService) *MessagesController {
	return &MessagesController{messageService: messageService}
}

func (ctl *MessagesController) GetAllMessages(c echo.Context) error {
	offsetStr := c.QueryParam("offset")
	offset := 0
	if offsetStr != "" {
		if parsed, err := strconv.Atoi(offsetStr); err == nil && parsed >= 0 {
			offset = parsed
		} else {
			return c.JSON(http.StatusBadRequest, dto.NewErrorResponse("invalid offset param"))
		}
	}

	limitStr := c.QueryParam("limit")
	limit := 20
	if limitStr != "" {
		if parsed, err := strconv.Atoi(limitStr); err == nil && parsed > 0 {
			limit = parsed
		} else {
			return c.JSON(http.StatusBadRequest, dto.NewErrorResponse("invalid limit param"))
		}
	}

	messages, err := ctl.messageService.GetAllMessages(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, dto.NewDataResponse(messages))
}
