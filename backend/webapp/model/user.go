package model

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Phone    string `form:"phone" json:"phone" binding:"required"`
}

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
