package controller

import (
	"github.com/gin-gonic/gin"
	"gorm/database"
	"gorm/models"
	"net/http"
)

func HandleCreate(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func HandleMultipleCreate(ctx *gin.Context) {

	var users []models.User
	err := ctx.ShouldBindJSON(&users)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	result := database.DB.Create(&users)
	if result.Error != nil {
		ctx.JSON(400, result.Error)
		return
	}

	ctx.JSON(200, "success")
}
