package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Children struct {
	gorm.Model
	ChildName      string    `json:"child_name"`
	ChildAge       int       `json:"child_age"`
	ChildBirthDate time.Time `json:"time"`
}
