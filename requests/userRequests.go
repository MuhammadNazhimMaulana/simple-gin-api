package requests

import (
	"simple-api/models"
)

// For Create User
type RequestCreateUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
	StatusUser  string
	Gender      string
}

// For Getting All user
type ResponseGetUser struct {
	Users        []models.User `json:"users"`
	ResponseCode string        `json:"responseCode"`
	Status       string        `json:"status"`
	Total        int64         `json:"total"`
}

// For Creating A User
type ResponseCreateUser struct {
	Data         models.User `json:"data"`
	ResponseCode string      `json:"responseCode"`
	Status       string      `json:"status"`
	Total        int64       `json:"total"`
}

// For Find One user
type ResponseFindUser struct {
	User         models.User `json:"user"`
	ResponseCode string      `json:"responseCode"`
	Status       string      `json:"status"`
	Total        int64       `json:"total"`
}
