package controller

import (
	"github.com/gin-gonic/gin"
	"gorm/database"
	"gorm/models"
)

func HandleGetUser(ctx *gin.Context) {

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

	var user *models.User
	result := database.DB.First(&user, "ID = ?", body.ID)
	if result.Error != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(200, gin.H{
		"user":    user,
		"message": "success",
	})
}
