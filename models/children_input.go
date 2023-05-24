package models

import "time"

type CreateChildrenInput struct {
	ChildName      string    `json:"child_name"`
	ChildAge       int       `json:"child_age"`
	ChildBirthDate time.Time `json:"time"`
	UserID         uint      `json:"user_id" binding:"required"`
}

type UpdateChildrenInput struct {
	ChildName      string    `json:"child_name"`
	ChildAge       int       `json:"child_age"`
	ChildBirthDate time.Time `json:"time"`
	UserID         uint      `json:"user_id" binding:"required"`
}
