package interfaces

import (
	"simple-api/models"
	"simple-api/requests"
)

type UserInterface interface {
	FindAll() (requests.ResponseGetUser, error)
	Find(id int) (requests.ResponseFindUser, error)
	Create(req models.User) (requests.ResponseCreateUser, error)
}
