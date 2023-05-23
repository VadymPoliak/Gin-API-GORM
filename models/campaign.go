package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Campaign struct {
	gorm.Model
	Name      string    `json:"name"`
	BeginDate time.Time `json:"time"`
	Events    []Event   `gorm:"foreignKey:Id;references:ID"`
}
