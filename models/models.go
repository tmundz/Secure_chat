package models

import (
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	Username string `json: "username" gorm:"test;not null;default:null"`
	Password string `json: "password" gorm:"test;not null;default:null:"`
}
