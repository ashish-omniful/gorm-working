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
	r.POST("/mapCreate", controller.HandleCreateMap)

	r.GET("/get", controller.HandleGetUser)
	r.GET("/firstRow", controller.HandleGetFirstRow)
	r.GET("/lastRow", controller.HandleGetLastRow)
	r.GET("/getAll", controller.HandlerGetAll)
	r.GET("/getByCondition", controller.HandleGetRecordByCondition)
	r.GET("/getByIN", controller.HandleGetRecordByIN)
	r.GET("/getByAND", controller.HandleGetRecordByAnd)
	r.GET("/getByNOT", controller.HandlerNotCondition)
	r.GET("getByOR", controller.HandlerOrCondition)

	r.DELETE("/deleteByPrimaryKey", controller.HandleDeleteByPrimary)

	r.PUT("updateRecord", controller.HandleUpdateRecord)
	r.PUT("/upsert", controller.HandleOnConflict)

	err := r.Run(":8080")
	if err != nil {
		panic("connection not established at port 8080")
	}
}
