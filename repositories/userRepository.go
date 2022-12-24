package repositories

import (
	"simple-api/models"
	"simple-api/requests"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo UserRepository) FindAll() (requests.ResponseGetUser, error) {
	users := []models.User{}
	repo.DB.Find(&users)

	// Response
	return requests.ResponseGetUser{
		Users:        users,
		ResponseCode: "00",
		Status:       "Success",
		Total:        int64(len(users)),
	}, nil
}
