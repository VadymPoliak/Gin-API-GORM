package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"project/db"
	"project/models"
)

func GetUsers(context *gin.Context) {

	var users []models.User

	db.DB.Find(&users)

	context.JSON(http.StatusOK, gin.H{"data": users})

}

func GetUser(context *gin.Context) {
	var user models.User

	if err := db.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
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
