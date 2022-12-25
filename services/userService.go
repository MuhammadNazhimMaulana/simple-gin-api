package services

import (
	"fmt"
	"simple-api/models"
	"simple-api/repositories/interfaces"
	"simple-api/requests"
	"time"
)

// For All User
func UserAll(repo interfaces.UserInterface) (requests.ResponseGetUser, error) {
	data, err := repo.FindAll()
	return data, err
}

// For Finding User
func User(repo interfaces.UserInterface, id int) (requests.ResponseFindUser, error) {
	data, err := repo.Find(id)
	return data, err
}

// For Creating User
func CreateUser(repo interfaces.UserInterface, req requests.RequestCreateUser) (requests.ResponseCreateUser, error) {

	date, err := time.Parse("DD-MM-YYYY", req.DateOfBirth)

	// If There Is An Error
	if err != nil {
		fmt.Println(err)
	}

	// Prepare Sent data
	user := models.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Address:     req.Address,
		DateOfBirth: date,
		StatusUser:  req.StatusUser,
		Gender:      req.Gender,
	}

	data, err := repo.Create(user)
	return data, err
}
