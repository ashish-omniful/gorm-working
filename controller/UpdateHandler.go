package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// where condition is necessary with &models.User{}
// use .Model(&user) to update that column
func HandleUpdateSingleColumn(ctx *gin.Context) {

	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	result := database.DB.Model(&user).Update("name", "changed")
	if result.Error != nil {
		ctx.JSON(400, result.Error)
		return
	}

	ctx.JSON(200, "SUCCESS")
	return
}

// when using map interfaces , use json column name
// when using User model , use struct variable name
func HandleUpdateSelectedColumn(ctx *gin.Context) {

	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
	result := database.DB.Model(&user).Omit("name").Updates(map[string]interface{}{
		"name":     "changed again",
		"password": "hanbhai",
	})

	if result.Error != nil {
		ctx.JSON(400, result.Error)
		return
	}

	ctx.JSON(200, "SUCCESS")
}

// update column for all rows
func HandleBatchUpdate(ctx *gin.Context) {

	result := database.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&models.User{}).Update("name", "yo")
	if result.Error != nil {
		ctx.JSON(400, result.Error)
		return
	}

	ctx.JSON(200, "SUCCESS")
}
