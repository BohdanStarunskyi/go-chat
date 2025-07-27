package services

import (
	"chat/database"
	"chat/models"
)

func GetUser(id int64) (models.User, error) {
	var user models.User
	result := database.DB.Where("id = ?", id).First(&user)
	return user, result.Error
}
