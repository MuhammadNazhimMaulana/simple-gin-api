package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	// Rendering Html
	ctx.HTML(http.StatusOK, "Home", gin.H{
		"content": "Halaman Index",
		"title":   "Home",
	})

}
