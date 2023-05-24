package controllers

import (
	"fmt"
	"net/http"
	"project/db"
	"project/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {

	var users []models.User

	// db.DB.Find(&users)
	db.DB.Preload("Children").Preload("Campaigns").Find(&users)

	context.JSON(http.StatusOK, gin.H{"data": users})

}

func GetUser(context *gin.Context) {
	var user models.User
	id := context.Param("id")

	if err := db.DB.Preload("Children").Preload("Campaigns").First(&user, id).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(context *gin.Context) {

	var input models.CreateUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Phone:     input.Phone,
		Password:  input.Password,
		Children:  input.Children,
		Campaigns: input.Campaigns,
	}

	fmt.Println("user", user)

	var a = db.DB.Create(&user)

	fmt.Print("user", a)

	context.JSON(http.StatusOK, gin.H{"data": user})

}

func UpdateUser(context *gin.Context) {
	var user models.User
	if err := db.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	var input models.UpdateUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&user).Updates(input)

	context.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(context *gin.Context) {
	var user models.User
	if err := db.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	db.DB.Delete(&user)

	context.JSON(http.StatusOK, gin.H{"data": true})
}

func TestCreateUser(context *gin.Context, user models.User) {

	db.DB.Create(&user)

	context.JSON(http.StatusOK, gin.H{"data": user})

}
