package services

import (
	"chat/database"
	"chat/dto"
	"chat/models"
	"chat/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type AuthService interface {
	Login(req dto.LoginRequest) (dto.AuthResponse, error)
	SignUp(req dto.SignupRequest) (dto.AuthResponse, error)
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) Login(req dto.LoginRequest) (dto.AuthResponse, error) {

	var user models.User
	result := database.DB.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		return dto.AuthResponse{}, result.Error
	}
	err := utils.ValidatePassword(req.Password, user.Password)

	if err != nil {
		return dto.AuthResponse{}, err
	}

	token, err := utils.GetJwtToken(user.ID)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	response := dto.AuthResponse{
		Token: token,
		User:  user.ToUserResponse(),
	}
	return response, nil
}

func (s *authService) SignUp(req dto.SignupRequest) (dto.AuthResponse, error) {
	var existing models.User
	if err := database.DB.Where("email = ?", req.Email).First(&existing).Error; err == nil {
		return dto.AuthResponse{}, fmt.Errorf("email already in use")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.AuthResponse{}, err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	user := models.NewUser(req.Name, req.Email, hashedPassword)
	if result := database.DB.Create(&user); result.Error != nil {
		return dto.AuthResponse{}, result.Error
	}

	token, err := utils.GetJwtToken(user.ID)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	return dto.AuthResponse{
		Token: token,
		User:  user.ToUserResponse(),
	}, nil
}
