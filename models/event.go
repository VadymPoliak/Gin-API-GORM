package models

type Event struct {
	Id     uint   `gorm:"primary_key;auto_increment" json:"id"`
	UserId User   `gorm:"references:Id"`
	Name   string `json:"first_name"`
}
