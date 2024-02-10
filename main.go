package main

import (
	"github.com/akshatsonic/go-server-learning/config"
	"github.com/akshatsonic/go-server-learning/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router) 
	router.Run(":8080")
}