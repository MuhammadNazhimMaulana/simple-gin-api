package main

import (
	"simple-api/configs"
	"simple-api/models"
	"simple-api/routers"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	db = configs.ConnectToMysql()
	db.AutoMigrate(&models.User{})
}

func main() {
	var PORT = ":8081"

	routers.StartServer().Run(PORT)
}
