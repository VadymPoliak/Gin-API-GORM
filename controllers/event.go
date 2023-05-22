package controllers

import (
	"fmt"
	"net/http"
	"project/db"
	"project/models"

	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {

	var users []models.Event

	db.DB.Find(&users)

	context.JSON(http.StatusOK, gin.H{"data": users})

}

func GetEvent(context *gin.Context) {
	var user models.Event

	if err := db.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateEvent(context *gin.Context) {

	var input models.CreateEventInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event := models.Event{
		Name: input.Name,
	}

	fmt.Println("event", event)

	var a = db.DB.Create(&event)

	fmt.Print("user", a)

	context.JSON(http.StatusOK, gin.H{"data": event})

}

func UpdateEvent(context *gin.Context) {
	var event models.Event
	if err := db.DB.Where("id = ?", context.Param("id")).First(&event).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	var input models.UpdateUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&event).Updates(input)

	context.JSON(http.StatusOK, gin.H{"data": event})
}

func DeleteEvent(context *gin.Context) {
	var event models.Event
	if err := db.DB.Where("id = ?", context.Param("id")).First(&event).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	db.DB.Delete(&event)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
