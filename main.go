package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm/controller"
	"gorm/database"
)

var DB *gorm.DB

func main() {

	fmt.Println("check")

	database.Init()

	r := gin.Default()
	r.POST("/create", controller.HandleCreate)
	r.GET("/get", controller.HandleGetUser)

	err := r.Run(":8080")
	if err != nil {
		panic("connection not established at port 8080")
	}
}
