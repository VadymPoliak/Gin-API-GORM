package controllers

import (
	"fmt"
	"net/http"

	"project/db"
	"project/models"

	"github.com/gin-gonic/gin"
)

func GetCampaignByUser(context *gin.Context) {

	var user models.User
	id := context.Param("user_id")

	db.DB.Preload("Campaigns").First(&user, id)

	context.JSON(http.StatusOK, gin.H{"data": user.Campaigns})
}

func CreateCampaign(context *gin.Context) {

	var input models.CreateCampaignInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	campaign := models.Campaign{
		Name:      input.Name,
		BeginDate: input.BeginDate,
		Events:    input.Events,
		UserID:    input.UserID,
	}

	fmt.Println("campaign", campaign)

	var a = db.DB.Create(&campaign)

	fmt.Print("campaign", a)

	context.JSON(http.StatusOK, gin.H{"data": campaign})

}

func UpdateCampaign(context *gin.Context) {
	var campaign models.Campaign
	if err := db.DB.Where("id = ?", context.Param("id")).First(&campaign).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Campaign not found!"})
		return
	}

	var input models.UpdateCampaignInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&campaign).Updates(input)

	context.JSON(http.StatusOK, gin.H{"data": campaign})
}

func DeleteCampaign(context *gin.Context) {
	var campaign models.Campaign
	if err := db.DB.Where("id = ?", context.Param("id")).First(&campaign).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Campaign not found!"})
		return
	}

	db.DB.Delete(&campaign)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
