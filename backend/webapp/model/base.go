package model

import (
	"gorm.io/gorm"
)

type BaseReview struct {
	gorm.Model
	ReviewTitle string `form:"ReviewTitle" json:"reviewtitle" binding:"required"`
	Review      string `form:"Review" json:"review" binding:"required"`
	Rating      int    `form:"Rating" json:"rating" binding:"required"`
	PlaceID     int    `form:"PlaceID" json:"placeid" binding:"required"`
	ReviewerID  int    `form:"ReviewerID" json:"reviewerid"`
}

type Places struct {
	gorm.Model
	Placename string `form:"Placename" json:"placename" binding:"required"`
	Location  string `form:"Location" json:"location" binding:"required"`
	Type      string `form:"Type" json:"type" binding:"required"`
	AvgRating int    `form:"AvgRating" json:"avgrating" binding:"required"`
}
