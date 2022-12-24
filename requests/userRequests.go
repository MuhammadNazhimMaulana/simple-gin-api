package requests

import "simple-api/models"

// For Getting All user
type ResponseGetUser struct {
	Users        []models.User `json:"users"`
	ResponseCode string        `json:"responseCode"`
	Status       string        `json:"status"`
	Total        int64         `json:"total"`
}
