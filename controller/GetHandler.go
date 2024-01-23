package controller

import (
	"github.com/gin-gonic/gin"
	"gorm/database"
	"gorm/models"
)

func HandleGetUser(ctx *gin.Context) {

	var body struct {
		ID uint `json:"id"`
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

func HandleGetFirstRow(ctx *gin.Context) {

	var user models.User

	result := database.DB.First(&user)
	if result.Error != nil {
		ctx.JSON(400, result.Error)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    user,
	})
}

func HandleGetLastRow(ctx *gin.Context) {
	var user models.User

	result := database.DB.Last(&user)
	if result.Error != nil {
		ctx.JSON(400, result.Error)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    user,
	})
}

func HandlerGetAll(ctx *gin.Context) {
	var users []models.User

	result := database.DB.Find(&users)
	if result.Error != nil {
		ctx.JSON(400, result.Error)
		return
	}

	ctx.JSON(200, users)
}

func HandleGetRecordByCondition(ctx *gin.Context) {

	var user models.User

	var body struct {
		Name string `json:"name"`
	}

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	result := database.DB.Where("name = ?", body.Name).First(&user) // can get all records corresponding to name using Find
	if result.Error != nil {
		ctx.JSON(400, result.Error)
		return
	}

	ctx.JSON(200, user)
}

func HandleGetRecordByIN(ctx *gin.Context) {

	var users []models.User
	result := database.DB.Where("name IN ?", []string{"user4", "user5"}).Find(&users)
	// result := database.DB.Where("name LIKE ? ", "user%").Find(&users)
	if result.Error != nil {
		ctx.JSON(400, result.Error)
		return
	}

	ctx.JSON(200, users)
}

func HandleGetRecordByAnd(ctx *gin.Context) {

	var users []models.User
	result := database.DB.Where("name = ? AND id >= ?", "ashish", 10).Find(&users)
	// for BETWEEN -> "name BETWEEN ? AND ?"
	if result.Error != nil {
		ctx.JSON(400, result.Error)
		return
	}

	ctx.JSON(200, users)
}
