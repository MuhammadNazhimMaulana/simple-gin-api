package routers

import (
	"simple-api/controllers"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// Configuration For Template
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})

	// Define Asset
	router.Static("/public", "./public")

	// Loading Html
	router.LoadHTMLGlob("views/**/*")

	// Home
	router.GET("/home", controllers.Index)

	// User
	router.GET("/user", controllers.AllUser)
	router.GET("/user/:userID", controllers.FindUser)
	router.POST("/user", controllers.CreateUser)

	return router
}
