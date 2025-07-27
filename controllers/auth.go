package controllers

import (
	"chat/dto"
	"chat/services"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ctl *AuthController) Login(c echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
	}

	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
	}

	res, err := ctl.authService.Login(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, dto.NewDataResponse(res))
}

func (ctl *AuthController) Signup(c echo.Context) error {
	var req dto.SignupRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
	}

	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
	}

	res, err := ctl.authService.SignUp(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, dto.NewDataResponse(res))
}
