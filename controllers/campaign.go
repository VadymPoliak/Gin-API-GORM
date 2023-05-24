package controllers

import (
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
