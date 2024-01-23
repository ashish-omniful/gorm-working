package controller

import (
	"github.com/gin-gonic/gin"
	"gorm/database"
	"gorm/models"
)

func HandleDeleteByPrimary(ctx *gin.Context) {

	var body struct {
		ID string `json:"id"`
	}

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	result := database.DB.Delete(&models.User{}, body.ID)
	if result.Error != nil {
		ctx.JSON(400, gin.H{
			"error": result.Error,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success",
	})
}
