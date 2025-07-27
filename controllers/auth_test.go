package controllers_test

import (
	"bytes"
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

func newEchoContext(t *testing.T, method, path string, payload any) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	req := httptest.NewRequest(method, path, bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func unmarshalAuthResponse(t *testing.T, rec *httptest.ResponseRecorder) dto.DefaultResponseWrapper[dto.AuthResponse] {
	var resp dto.DefaultResponseWrapper[dto.AuthResponse]
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	return resp
}

func TestAuthController_Login_Success(t *testing.T) {
	mockAuthService := &mock.MockAuthService{
		LoginFunc: func(req dto.LoginRequest) (dto.AuthResponse, error) {
			return dto.AuthResponse{
				Token: "mock-token",
				User: dto.UserResponse{
					ID:   1,
					Name: "Mock User",
				},
			}, nil
		},
	}
	controller := controllers.NewAuthController(mockAuthService)

	loginPayload := dto.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	c, rec := newEchoContext(t, http.MethodPost, "/login", loginPayload)

	err := controller.Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	resp := unmarshalAuthResponse(t, rec)
	assert.Equal(t, "mock-token", resp.Data.Token)
}

func TestAuthController_Signup_Success(t *testing.T) {
	mockAuthService := &mock.MockAuthService{
		SignUpFunc: func(req dto.SignupRequest) (dto.AuthResponse, error) {
			return dto.AuthResponse{
				Token: "mock-token",
				User: dto.UserResponse{
					ID:   1,
					Name: req.Name,
				},
			}, nil
		},
	}
	controller := controllers.NewAuthController(mockAuthService)

	signupPayload := dto.SignupRequest{
		Email:    "test@example.com",
		Password: "password123",
		Name:     "TestName",
	}
	c, rec := newEchoContext(t, http.MethodPost, "/signup", signupPayload)

	err := controller.Signup(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	resp := unmarshalAuthResponse(t, rec)
	assert.Equal(t, "mock-token", resp.Data.Token)
	assert.Equal(t, signupPayload.Name, resp.Data.User.Name)
}
