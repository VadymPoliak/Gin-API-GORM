package models

type CreateEventInput struct {
	Id     uint   `gorm:"primary_key;auto_increment" json:"id"`
	UserId User   `gorm:"references:Id"`
	Name   string `json:"first_name"`
}

type UpdateEventInput struct {
	Id     uint   `gorm:"primary_key;auto_increment" json:"id"`
	UserId User   `gorm:"references:Id"`
	Name   string `json:"first_name"`
}
