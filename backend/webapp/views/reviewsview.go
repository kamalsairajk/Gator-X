package views

import (
	"fmt"
	"net/http"
	"strconv"

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
		if len(reviews) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "currently the database is empty",
				"result":  reviews,
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
		fmt.Println(json)

		p := bluemonday.StripTagsPolicy()

		json.ReviewTitle = p.Sanitize(json.ReviewTitle)
		json.Review = p.Sanitize(json.Review)
		json.Rating, _ = strconv.Atoi(p.Sanitize(strconv.Itoa(json.Rating)))

		json.PlaceID, _ = strconv.Atoi(p.Sanitize(strconv.Itoa(json.PlaceID)))
		json.ReviewerID, _ = strconv.Atoi(p.Sanitize(strconv.Itoa(json.ReviewerID)))

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

		reviewid, _ := strconv.Atoi(c.Query("reviewID"))
		reviewerid, _ := strconv.Atoi(c.Query("reviewerID"))

		var breview model.BaseReview
		result := db.Model(&breview).Where("id = ? AND reviewer_id = ?", reviewid, reviewerid).Updates(&json)

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
		// var json model.BaseReview
		// if err := c.ShouldBindJSON(&json); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		reviewid, _ := strconv.Atoi(c.Query("reviewID"))
		reviewerid, _ := strconv.Atoi(c.Query("reviewerID"))
		fmt.Println(reviewerid, reviewid)

		var breview model.BaseReview
		db.Find(&breview, "id = ? AND reviewer_id = ?", reviewid, reviewerid)
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
