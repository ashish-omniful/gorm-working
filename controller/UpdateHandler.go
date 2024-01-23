package controller

import (
	"github.com/gin-gonic/gin"
	"gorm/database"
	"gorm/models"
)

func HandleUpdateRecord(ctx *gin.Context) {

	var user *models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	result := database.DB.Save(&user) // updates if id > 0,  otherwise inserts
	if result.Error != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success",
	})
}
