package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"gorm/database"
	"gorm/models"
)

func HandleOnConflict(ctx *gin.Context) {

	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	result := database.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // yahan table mein column ka naam
		UpdateAll: true,                          // can specify selected columns using doUpdates
	}).Create(&user)
	if result.Error != nil {
		ctx.JSON(400, result.Error)
		return
	}

	ctx.JSON(200, "success")
}
