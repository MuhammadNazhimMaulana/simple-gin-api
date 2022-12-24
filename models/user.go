package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Address     string    `json:"address"`
	DateOfBirth time.Time `json:"date_of_birth"`
	StatusUser  string
	Gender      string
	gorm.Model
}

type Table interface {
	TableName() string
}

func (user User) TableName() string {
	return "users"
}
