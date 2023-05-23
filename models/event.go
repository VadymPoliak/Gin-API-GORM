package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	Name       string      `json:"name"`
	Date       time.Time   `json:"date"`
	EventItems []EventItem `gorm:"foreignKey:Id;references:ID"`
}
