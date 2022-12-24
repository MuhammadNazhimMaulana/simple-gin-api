package services

import (
	"simple-api/repositories/interfaces"
	"simple-api/requests"
)

func UserAll(repo interfaces.UserInterface) (requests.ResponseGetUser, error) {
	data, err := repo.FindAll()
	return data, err
}
