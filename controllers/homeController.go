package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	ctx.JSON(http.StatusCreated, gin.H{
		"Test": "test",
	})

}
