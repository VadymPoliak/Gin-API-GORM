package controllers

import (
	"fmt"
	"net/http"
	"project/db"
	"project/models"

	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {

	var events []models.Event

	db.DB.Preload("EventItems").Find(&events)

	context.JSON(http.StatusOK, gin.H{"data": events})

}

func GetEventByID(context *gin.Context) {
	var event models.Event
	id := context.Param("id")

	if err := db.DB.Preload("EventItems").First(&event, id).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Event not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": event})
}

func CreateEvent(context *gin.Context) {

	var input models.CreateEventInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event := models.Event{
		Name:       input.Name,
		Date:       input.Date,
		CampaignID: input.CampaignID,
		EventItems: input.EventItems,
	}

	fmt.Println("event", event)

	var a = db.DB.Create(&event)

	fmt.Print("event", a)

	context.JSON(http.StatusOK, gin.H{"data": event})

}

func UpdateEvent(context *gin.Context) {
	var event models.Event
	if err := db.DB.Where("id = ?", context.Param("id")).First(&event).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Event not found!"})
		return
	}

	var input models.UpdateEventInput
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
		context.JSON(http.StatusBadRequest, gin.H{"error": "Event not found!"})
		return
	}

	db.DB.Delete(&event)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
