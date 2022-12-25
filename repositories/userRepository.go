package repositories

import (
	"simple-api/models"
	"simple-api/requests"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// Find All
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

// Find One
func (repo UserRepository) Find(id int) (requests.ResponseFindUser, error) {
	user := models.User{}
	repo.DB.Find(&user, id)

	// Response
	return requests.ResponseFindUser{
		User:         user,
		ResponseCode: "00",
		Status:       "Success",
		Total:        1,
	}, nil
}

// Create One
func (repo UserRepository) Create(user models.User) (requests.ResponseCreateUser, error) {
	repo.DB.Create(&user)

	// Response
	return requests.ResponseCreateUser{
		Data:         user,
		ResponseCode: "00",
		Status:       "Success",
		Total:        1,
	}, nil
}
