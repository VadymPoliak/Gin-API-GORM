package models

import "time"

type CreateCampaignInput struct {
	Name      string    `json:"name"`
	BeginDate time.Time `json:"time"`
	Events    []Event   `gorm:"foreignKey:CampaignID"`
	UserID    uint      `json:"user_id" binding:"required"`
}

type UpdateCampaignInput struct {
	Name      string    `json:"name"`
	BeginDate time.Time `json:"time"`
	Events    []Event   `gorm:"foreignKey:CampaignID"`
	UserID    uint      `json:"user_id" binding:"required"`
}
