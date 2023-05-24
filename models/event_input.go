package models

import "time"

type CreateEventInput struct {
	Name       string      `json:"name"`
	Date       time.Time   `json:"date"`
	CampaignID uint        `json:"campaign_id" binding:"required"`
	EventItems []EventItem `gorm:"foreignKey:EventID"`
}

type UpdateEventInput struct {
	Name       string      `json:"name"`
	Date       time.Time   `json:"date"`
	CampaignID uint        `json:"campaign_id" binding:"required"`
	EventItems []EventItem `gorm:"foreignKey:EventID"`
}
