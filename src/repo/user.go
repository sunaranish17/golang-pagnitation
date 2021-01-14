package repo

import (
	"golang-pagination/src/config"
	models "golang-pagination/src/models"
)

func GetAllUsers(user *models.User, pagination *models.Pagination) (*[]models.User, error) {
	var users []models.User
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuilder.Model(&models.User{}).Where(user).Find(&users)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &users, nil
}
