package controllers

import (
	"net/http"
	"simple-api/configs"
	"simple-api/repositories"
	"simple-api/requests"
	"simple-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = configs.ConnectToMysql()
}

func AllUser(ctx *gin.Context) {
	var res requests.ResponseGetUser

	repo := repositories.UserRepository{
		DB: db,
	}

	res, err := services.UserAll(repo)

	// Jika gagal
	if err != nil {
		res.Status = "Get All User Gagal"
		res.ResponseCode = "400"
	}

	ctx.JSON(http.StatusOK, res)
}
