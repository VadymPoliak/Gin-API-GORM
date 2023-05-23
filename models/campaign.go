package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Campaign struct {
	gorm.Model
	CampaignId int       `gorm:"primary_key" json:"campaign_id"`
	Name       string    `json:"name"`
	BeginDate  time.Time `json:"time"`
	Events     []Event   `gorm:"foreignKey:Id;references:CampaignId"`
}
