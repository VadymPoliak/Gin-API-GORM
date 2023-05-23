package controllers

import (
	"net/http"
	"time"

	"project/db"
	"project/models"

	"github.com/gin-gonic/gin"
)

func GetCampaign(context *gin.Context) {

	var user models.User

	db.DB.Preload("Campaigns").First(&user)

	context.JSON(http.StatusOK, gin.H{"data": user.Children})
}

func AddCampaign(context *gin.Context) {
	var user models.User

	beginDate, err := time.Parse("2006-01-02", "2015-10-08")
	if err != nil {
		panic(err)
	}

	user.Campaigns = append(user.Campaigns, models.Campaign{Name: "XXX", BeginDate: beginDate})
	db.DB.Save(&user)

	context.JSON(http.StatusOK, gin.H{"data": user})
}
