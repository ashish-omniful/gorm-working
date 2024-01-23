package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm/controller"
	"gorm/database"
)

func main() {

	fmt.Println("check")

	database.Init()

	r := gin.Default()
	r.POST("/create", controller.HandleCreate)
	r.POST("/multipleCreate", controller.HandleMultipleCreate)
	r.POST("/select-omit", controller.HandleSelectOmitCreate)
	r.GET("/get", controller.HandleGetUser)
	r.DELETE("/deleteByPrimaryKey", controller.HandleDeleteByPrimary)
	r.PUT("updateRecord", controller.HandleUpdateRecord)

	err := r.Run(":8080")
	if err != nil {
		panic("connection not established at port 8080")
	}
}
