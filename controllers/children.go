package controllers

import (
	"fmt"
	"net/http"

	"project/db"
	"project/models"

	"github.com/gin-gonic/gin"
)

func GetChildrenByUser(context *gin.Context) {

	var user models.User
	id := context.Param("user_id")

	db.DB.Preload("Children").First(&user, id)

	context.JSON(http.StatusOK, gin.H{"data": user.Children})
}

func CreateChildren(context *gin.Context) {

	var input models.CreateChildrenInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	children := models.Children{
		ChildName:      input.ChildName,
		ChildAge:       input.ChildAge,
		ChildBirthDate: input.ChildBirthDate,
		UserID:         input.UserID,
	}

	fmt.Println("children", children)

	var a = db.DB.Create(&children)

	fmt.Print("children", a)

	context.JSON(http.StatusOK, gin.H{"data": children})

}

func UpdateChildren(context *gin.Context) {
	var children models.Children
	if err := db.DB.Where("id = ?", context.Param("id")).First(&children).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Children not found!"})
		return
	}

	var input models.UpdateChildrenInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&children).Updates(input)

	context.JSON(http.StatusOK, gin.H{"data": children})
}

func DeleteChildren(context *gin.Context) {
	var children models.Children
	if err := db.DB.Where("id = ?", context.Param("id")).First(&children).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Children not found!"})
		return
	}

	db.DB.Delete(&children)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
