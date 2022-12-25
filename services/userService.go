package services

import (
	"simple-api/repositories/interfaces"
	"simple-api/requests"
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
