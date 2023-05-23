package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	Name   string       `json:"name"`
	Date   time.Time    `json:"date"`
	Events []EventItems `gorm:"foreignKey:ID;references:ID"`
}
