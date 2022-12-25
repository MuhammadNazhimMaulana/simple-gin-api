package interfaces

import (
	"simple-api/requests"
)

type UserInterface interface {
	FindAll() (requests.ResponseGetUser, error)
	Find(id int) (requests.ResponseFindUser, error)
}
