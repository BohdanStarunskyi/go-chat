package mock

import "chat/dto"

type MockAuthService struct {
	LoginFunc  func(dto.LoginRequest) (dto.AuthResponse, error)
	SignUpFunc func(dto.SignupRequest) (dto.AuthResponse, error)
}

func (m *MockAuthService) Login(req dto.LoginRequest) (dto.AuthResponse, error) {
	return m.LoginFunc(req)
}

func (m *MockAuthService) SignUp(req dto.SignupRequest) (dto.AuthResponse, error) {
	return m.SignUpFunc(req)
}
