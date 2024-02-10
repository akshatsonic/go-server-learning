package controller

import (
	"fmt"

	"github.com/akshatsonic/go-server-learning/config"
	"github.com/akshatsonic/go-server-learning/models"
	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context){
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func GetAllUsers(c *gin.Context){
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func AddUser(c *gin.Context){
	var requestUser models.User
	if err := c.BindJSON(&requestUser);err!=nil{
		panic("error in request body")
	}
	config.DB.Save(&requestUser)
	fmt.Println(requestUser.Email)
}