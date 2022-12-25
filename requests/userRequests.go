package requests

import "simple-api/models"

// For Getting All user
type ResponseGetUser struct {
	Users        []models.User `json:"users"`
	ResponseCode string        `json:"responseCode"`
	Status       string        `json:"status"`
	Total        int64         `json:"total"`
}

// For Find One user
type ResponseFindUser struct {
	User         models.User `json:"user"`
	ResponseCode string      `json:"responseCode"`
	Status       string      `json:"status"`
	Total        int64       `json:"total"`
}
