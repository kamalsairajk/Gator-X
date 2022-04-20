package model

import (
	"gorm.io/gorm"
)

type UserType int
// different users
const (
	NORMAL   UserType = iota
	ADMIN             //status =1
)

// User table
type Users struct {
	gorm.Model
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Type	UserType

}

// Login struct
type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
