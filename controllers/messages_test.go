package controllers_test

import (
	"chat/controllers"
	"chat/dto"
	"chat/mock"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupTestContext(query string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/messages"+query, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func TestGetAllMessages_Success(t *testing.T) {
	mockService := &mock.MockMessageService{}
	controller := controllers.NewMessageController(mockService)

	c, rec := setupTestContext("?offset=0&limit=10")

	err := controller.GetAllMessages(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response dto.DefaultResponseWrapper[[]dto.MessageResponse]
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response.Data, 1)
	assert.Equal(t, "Test message", response.Data[0].Message)
	assert.Equal(t, int64(1), response.Data[0].Sender.ID)
}
