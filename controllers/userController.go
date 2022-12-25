package controllers

import (
	"net/http"
	"simple-api/configs"
	"simple-api/repositories"
	"simple-api/requests"
	"simple-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = configs.ConnectToMysql()
}

func AllUser(ctx *gin.Context) {
	var res requests.ResponseGetUser

	// Instance of Repo
	repo := repositories.UserRepository{
		DB: db,
	}

	// Getting Result
	res, err := services.UserAll(repo)

	// If Fail
	if err != nil {
		res.Status = "Get All User Failed"
		res.ResponseCode = "400"
	}

	ctx.JSON(http.StatusOK, res)
}

func FindUser(ctx *gin.Context) {
	var res requests.ResponseFindUser

	// Getting id
	id, errstr := strconv.Atoi(ctx.Param("userID"))

	if errstr != nil {
		res.Status = "No ID Included"
		res.ResponseCode = "400"
	}

	repo := repositories.UserRepository{
		DB: db,
	}

	res, err := services.User(repo, id)

	// Jika gagal
	if err != nil {
		res.Status = "Get User Failed"
		res.ResponseCode = "400"
	}

	ctx.JSON(http.StatusOK, res)
}

func CreateUser(ctx *gin.Context) {
	var req requests.RequestCreateUser
	var res requests.ResponseCreateUser

	// Check Error
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	repo := repositories.UserRepository{
		DB: db,
	}

	res, err := services.CreateUser(repo, req)

	// Jika gagal
	if err != nil {
		res.Status = "Create User Failed"
		res.ResponseCode = "400"
	}

	ctx.JSON(http.StatusOK, res)
}
