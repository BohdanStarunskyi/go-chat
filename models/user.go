package models

import "chat/dto"

type User struct {
	ID       int64 `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

func (u User) ToUserResponse() dto.UserResponse {
	return dto.NewUserResponse(u.ID, u.Name)
}

func NewUser(
	name string,
	email string,
	password string,
) User {
	return User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
