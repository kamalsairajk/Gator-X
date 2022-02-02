package views

import (
	"fmt"
	"net/http"
	"strconv"

	// "time"
	model "webapp/model"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
)

func GetallreviewsView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var reviews []model.BaseReview

		db.Find(&reviews)

		if len(reviews) > 0 {
			c.JSON(http.StatusOK, gin.H{
				"result": reviews,
			})
			return
		}
	}

	return gin.HandlerFunc(fn)
}

func PostreviewView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var json model.BaseReview
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(&json)
		result := db.Create(&json)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": "Review created in database",
		})
	}
	return gin.HandlerFunc(fn)
}

func EditreviewView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var json model.BaseReview
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(json)

		p := bluemonday.StripTagsPolicy()
		name := p.Sanitize(json.Name)
		placename := p.Sanitize(json.Placename)
		location := p.Sanitize(json.Location)

		var breview model.BaseReview
		db.First(&breview, "name = ? AND placename = ?  AND location = ? ", name, placename, location)

		breview.Review = p.Sanitize(json.Review)
		i, err := strconv.Atoi(p.Sanitize(strconv.Itoa(json.Rating)))
		fmt.Println(err)
		breview.Rating = i
		result := db.Save(&breview)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": "Review edited in database",
		})
	}
	return gin.HandlerFunc(fn)
}

func DeletereviewView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var json model.BaseReview
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(json)

		p := bluemonday.StripTagsPolicy()
		name := p.Sanitize(json.Name)
		placename := p.Sanitize(json.Placename)
		location := p.Sanitize(json.Location)

		var breview model.BaseReview
		db.First(&breview, "name = ? AND placename = ?  AND location = ? ", name, placename, location)

		// breview.Review = p.Sanitize(json.Review)
		// i, err := strconv.Atoi(p.Sanitize(strconv.Itoa(json.Rating)))
		// fmt.Println(err)
		// breview.Rating = i
		result := db.Delete(&breview)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": "Review deleted from database",
		})
	}
	return gin.HandlerFunc(fn)
}
