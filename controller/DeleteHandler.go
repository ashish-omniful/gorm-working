package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"gorm/database"
	"gorm/models"
)

func HandleDeleteByPrimary(ctx *gin.Context) {

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

// delete does a batch delete if no primary key is given
func HandleBatchDelete(ctx *gin.Context) {

	users := []models.User{{ID: 1}, {ID: 2}}
	result := database.DB.Delete(&users)
	if result.Error != nil {
		ctx.JSON(400, result.Error)
		return
	}

	ctx.JSON(200, "successfully done batch delete")
}

func HandleReturnDeletedColumns(ctx *gin.Context) {

	var users []models.User
	result := database.DB.Clauses(clause.Returning{}).Where("name = ?", "yo").Delete(&users)
	if result.Error != nil {
		ctx.JSON(400, result.Error)
		return
	}

	ctx.JSON(200, users)
}

// soft delete erases and returns the deleted rows
