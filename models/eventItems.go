package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type EventItems struct {
	gorm.Model
	Name  string    `json:"name"`
	Date  time.Time `json:"date"`
	Taken bool      `json:"taken"`
}
