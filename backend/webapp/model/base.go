package model

import "gorm.io/gorm"

type BaseReview struct {
	gorm.Model
	Name   string `form:"Name" json:"name" binding:"required"`
	Review string `form:"Review" json:"review" binding:"required"`
	Rating int    `form:"Rating" json:"rating" binding:"required"`
}

type Places struct {
	gorm.Model
	Placename string `form:"Placename" json:"placename" binding:"required"`
	Location  string `form:"Location" json:"location" binding:"required"`
	Type      string `form:"Type" json:"type" binding:"required"`
	Rating    int    `form:"Rating" json:"rating" binding:"required"`
}
