package sockets

import (
	"bytes"
	"chat/dto"
	"chat/services"
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {
	UserID         int64
	Conn           *websocket.Conn
	Send           chan []byte
	MessageService services.MessageService
}

func NewClient(conn *websocket.Conn, userID int64) *Client {
	return &Client{
		UserID:         userID,
		Conn:           conn,
		Send:           make(chan []byte),
		MessageService: services.NewMessageService(),
	}
}

func (c *Client) Read() {
	defer func() {
		HubInstance.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		var req dto.MessageRequest
		decoder := json.NewDecoder(bytes.NewReader(message))
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&req)
		if err != nil {
			c.sendErrorMessage(err)
			continue
		}

		processed, err := c.MessageService.HandleIncomingMessage(c.UserID, req)
		if err != nil {
			c.sendErrorMessage(err)
			continue
		}
		c.sendSuccessMessage(processed)
	}
}

func (c *Client) sendErrorMessage(err error) {
	errorResp := dto.NewErrorResponse(err.Error())
	jsonBytes, marshalErr := json.Marshal(errorResp)
	if marshalErr != nil {
		return
	}
	c.Send <- jsonBytes
}

func (c *Client) sendSuccessMessage(data any) {
	successResponse := dto.NewDataResponse(data)
	jsonBytes, marshalErr := json.Marshal(successResponse)
	if marshalErr != nil {
		return
	}
	HubInstance.Broadcast <- jsonBytes
}

func (c *Client) Write() {
	defer c.Conn.Close()

	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}
