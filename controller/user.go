package controller

import (
	"github.com/akshatsonic/go-server-learning/config"
	"github.com/akshatsonic/go-server-learning/models"
	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context){
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}