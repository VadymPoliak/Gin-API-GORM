package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string     `json:"first_name" binding:"required"`
	LastName  string     `json:"last_name" binding:"required"`
	Email     string     `json:"email" binding:"required"`
	Phone     string     `json:"phone" binding:"required"`
	Password  string     `json:"password" binding:"required"`
	Children  []Children `gorm:"foreignKey:UserID"`
	Campaigns []Campaign `gorm:"foreignKey:UserID"`
}
