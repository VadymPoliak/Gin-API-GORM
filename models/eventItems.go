package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type EventItem struct {
	gorm.Model
	Name  string    `json:"name"`
	Date  time.Time `json:"date"`
	Taken bool      `json:"taken"`
}
