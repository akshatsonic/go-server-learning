package config

import (
	"github.com/akshatsonic/go-server-learning/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db,err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{ })
	if err != nil{
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB=db
}