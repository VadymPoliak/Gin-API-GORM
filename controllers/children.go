package controllers

import (
	"net/http"
	"time"

	"project/db"
	"project/models"

	"github.com/gin-gonic/gin"
)

func GetChildren(context *gin.Context) {

	var user models.User

	db.DB.Preload("Children").First(&user)

	context.JSON(http.StatusOK, gin.H{"data": user.Children})
}

func AddChildren(context *gin.Context) {
	var user models.User

	birthTime, err := time.Parse("2006-01-02", "2020-03-10")
	if err != nil {
		panic(err)
	}

	user.Children = append(user.Children, models.Children{ChildName: "John", ChildAge: 3, ChildBirthDate: birthTime})
	db.DB.Save(&user)

	context.JSON(http.StatusOK, gin.H{"data": user})
}
