package views

import (
	"fmt"
	"net/http"
	"strconv"
	"math"
	model "webapp/model"

	"github.com/gin-contrib/sessions"
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
		session := sessions.Default(c)
		i := session.Get("userId")
		if i == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user not logged in"})
			return
		}
		userId := i.(uint)
		var user model.Users
		db.First(&user, "id = ?", userId)

		var json model.BaseReview
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		p := bluemonday.StripTagsPolicy()

		json.ReviewTitle = p.Sanitize(json.ReviewTitle)
		json.Review = p.Sanitize(json.Review)
		json.Rating, _ = strconv.Atoi(p.Sanitize(strconv.Itoa(json.Rating)))

		json.PlaceID, _ = strconv.Atoi(p.Sanitize(strconv.Itoa(json.PlaceID)))
		// json.ReviewerID, _ = strconv.Atoi(p.Sanitize(strconv.Itoa(json.ReviewerID)))
		json.ReviewerID = int(user.ID)
		result := db.Create(&json)

		// feature - calculating average review for a place
		var placereviews []model.BaseReview
		db.Find(&placereviews, "place_id = ?", json.PlaceID)
		var avgrating =float64(0.0)
		for i := 0; i < len(placereviews); i++ {
			avgrating+=float64(placereviews[i].Rating)
		}
		avgrating=avgrating/float64(len(placereviews))
		avgrating=math.Round(avgrating*100)/100
		var uplace model.Places
		result1 := db.Model(&uplace).Where("id = ?", json.PlaceID).Update("avg_rating",avgrating)

		if result1.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result1.Error.Error()})
			return
		}
		// 
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
		session := sessions.Default(c)
		i := session.Get("userId")
		if i == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user not logged in"})
			return
		}
		userId := i.(uint)
		var user model.Users
		db.First(&user, "id = ?", userId)

		var json model.BaseReview
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(json)

		reviewid, _ := strconv.Atoi(c.Param("reviewID"))
		// reviewerid, _ := strconv.Atoi(c.Query("reviewerID"))
		reviewerid := int(user.ID)

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
		session := sessions.Default(c)
		i := session.Get("userId")
		if i == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user not logged in"})
			return
		}
		userId := i.(uint)
		var user model.Users
		db.First(&user, "id = ?", userId)
		reviewid, _ := strconv.Atoi(c.Param("reviewID"))
		// reviewerid, _ := strconv.Atoi(c.Query("reviewerID"))
		reviewerid := int(userId)
		// fmt.Println(reviewerid, reviewid)

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

func GetreviewsbyuserView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		session := sessions.Default(c)
		i := session.Get("userId")
		if i == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user not logged in"})
			return
		}
		userId := i.(uint)
		var user model.Users
		db.First(&user, "id = ?", userId)

		var userreviews []model.BaseReview
		result := db.Find(&userreviews, "reviewer_id = ?", user.ID)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": userreviews,
		})
	}
	return gin.HandlerFunc(fn)

}

func GetreviewsbyplaceView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		placeID, _ := strconv.Atoi(c.Param("placeID"))

		var place model.Places
		result := db.First(&place, "id = ?", placeID)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		var placereviews []model.BaseReview
		result1 := db.Find(&placereviews, "place_id = ?", place.ID)

		if result1.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": placereviews,
		})

	}
	return gin.HandlerFunc(fn)
}