package routes

import (
	"github.com/akshatsonic/go-server-learning/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)
}