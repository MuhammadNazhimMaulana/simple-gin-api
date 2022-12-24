package routers

import (
	"simple-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// Home
	router.GET("/home", controllers.Index)

	// User
	router.GET("/user", controllers.AllUser)

	return router
}
