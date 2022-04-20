package model

import (
	"gorm.io/gorm"
)
//base review table, mapped to json variables
type BaseReview struct {
	gorm.Model
	ReviewTitle string `form:"ReviewTitle" json:"reviewtitle"`
	Review      string `form:"Review" json:"review"`
	Rating      int    `form:"Rating" json:"rating"`
	PlaceID     int    `form:"PlaceID" json:"placeid"`
	ReviewerID  int    `form:"ReviewerID" json:"reviewerid"`
	ReviewImage 		string `form:"file" json:"file"`
}

//places table, mapped to json variables
type Places struct {
	gorm.Model
	Placename string `form:"Placename" json:"placename"`
	Location  string `form:"Location" json:"location"`
	Type      string `form:"Type" json:"type"`
	AvgRating int    `form:"AvgRating" json:"avgrating"`
	PlaceImage 		string `form:"file" json:"file"`

}
