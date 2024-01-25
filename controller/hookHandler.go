package controller

import (
	"github.com/gin-gonic/gin"
	"gorm/database"
	"gorm/models"
)

func HandleHook(ctx *gin.Context) {

	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(400, result.Error)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
	})
}
