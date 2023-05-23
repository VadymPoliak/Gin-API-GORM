package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserId    uint       `gorm:"primary_key;auto_increment" json:"user_id"`
	FirstName string     `json:"first_name" binding:"required"`
	LastName  string     `json:"last_name" binding:"required"`
	Email     string     `json:"email" binding:"required"`
	Phone     string     `json:"phone" binding:"required"`
	Password  string     `json:"password" binding:"required"`
	Children  []Children `gorm:"foreignKey:ChildName;references:UserId"`
	Campaigns []Campaign `gorm:"foreignKey:CampaignId;references:UserId"`
}
