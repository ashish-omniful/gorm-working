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
		ctx.JSON(400, gin.H{
			"error": "while binding json in multiple create",
			"err":   err,
		})
		return
	}

	result := database.DB.Create(&users)
	if result.Error != nil {
		ctx.JSON(400, gin.H{
			"error": "while creating multiple users",
			"err":   result.Error,
		})
		return
	}

	ctx.JSON(200, "success")
}

func HandleSelectOmitCreate(ctx *gin.Context) {

	var user, user2 models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	user2 = user

	result := database.DB.Select([]string{"Email", "Password"}).Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	result2 := database.DB.Omit("Email", "Password").Create(&user2)
	if result2.Error != nil {
		ctx.JSON(400, result2.Error)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func HandleCreateMap(ctx *gin.Context) {

	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	result := database.DB.Model(&models.User{}).Create(map[string]interface{}{
		"Name":     user.Name,
		"Email":    user.Email,
		"Password": user.Password,
	})
	if result.Error != nil {
		ctx.JSON(400, result.Error)
		return
	}

	ctx.JSON(200, "success")
}
