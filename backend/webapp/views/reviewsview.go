package views

import (
	"net/http"
	"strconv"
	"math"
	model "webapp/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
	"path/filepath"
	"github.com/google/uuid"
	"encoding/json"
	"os"
	"fmt"

)



/*
	GetallreviewsView - return all reviews, using an array of base review model to store the results, find the database, and if the length of this array is greater than
	0 return the results and if length of results is 0 then return a message that the table is currently empty.
*/

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


/*
	PostreviewView - create a review for a place with given details, using the current session if user logged in if not then user should login and only logged in users have
	can create a review, all review related details are added into a json object under the form field data and image related to the review is added
	into the file field of the form data, after performing some data processing the data is stored into the database. After storing since the average rating for this place
	might change we calculate this again and update this in the places table. And finally we return a message saying that review is created or else we return 
	the error. 
*/

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

		file, err := c.FormFile("file")
		newFilepath:=""
		if err == nil {
			extension := filepath.Ext(file.Filename)
			newFileName := uuid.New().String() + extension
			newFilepath="C:/Users/kamal/Documents/SE project/Gator-X/backend/webapp/images/reviewimages/" + newFileName
			if err := c.SaveUploadedFile(file, newFilepath); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to save the file",
				})
			return
			}
		}
		data,_:=c.GetPostForm("data")

		var json1 model.BaseReview
		json.Unmarshal([]byte(data), &json1)
		// if err := c.ShouldBindJSON(&json); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		p := bluemonday.StripTagsPolicy()

		json1.ReviewTitle = p.Sanitize(json1.ReviewTitle)
		json1.Review = p.Sanitize(json1.Review)
		json1.Rating, _ = strconv.Atoi(p.Sanitize(strconv.Itoa(json1.Rating)))

		json1.PlaceID, _ = strconv.Atoi(p.Sanitize(strconv.Itoa(json1.PlaceID)))
		// json.ReviewerID, _ = strconv.Atoi(p.Sanitize(strconv.Itoa(json.ReviewerID)))
		json1.ReviewerID = int(user.ID)
		json1.ReviewImage=newFilepath
		result := db.Create(&json1)

		// feature - calculating average review for a place
		var placereviews []model.BaseReview
		db.Find(&placereviews, "place_id = ?", json1.PlaceID)
		var avgrating =float64(0.0)
		for i := 0; i < len(placereviews); i++ {
			avgrating+=float64(placereviews[i].Rating)
		}
		avgrating=avgrating/float64(len(placereviews))
		avgrating=math.Round(avgrating)
		// fmt.Println(avgrating)
		var uplace model.Places
		result1 := db.Model(&uplace).Where("id = ?", json1.PlaceID).Update("avg_rating",avgrating)

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

/*
	EditreviewView - edit a review with given constraints or fields using the review id, using the current session if user logged in if not then user should login and only logged in users have
	can edit an existing review. Taking the review id as a parameter to edit and find the review in the database, if the review is not present then return a status bad request
	else if finding the review returns an error then return the error message, if found then update the record in the database with the given details and if file is present 
	delete existing review image in the server and save the current file in the server and change this in the database along with other details. Also if the rating is present then, change the rating
	and calculate and update the average rating in the places table. Finally return an message that says the review is edited else the error.
*/

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

		reviewid, _ := strconv.Atoi(c.Param("reviewID"))
		// reviewerid, _ := strconv.Atoi(c.Query("reviewerID"))
		reviewerid := int(user.ID)
		var breview model.BaseReview

		result1:=db.Find(&breview, "id = ? AND reviewer_id = ?", reviewid, reviewerid)
		if (breview==model.BaseReview{}){
			c.JSON(http.StatusBadRequest, gin.H{"error": "Review does not exist"})
			return
		}

		if result1.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result1.Error.Error()})
			return
		}
		// fmt.Println(place)

		file, err := c.FormFile("file")
		newFilepath:=""
		if breview.ReviewImage!="" && err==nil{
			e := os.Remove(breview.ReviewImage)
			if e != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to delete the file",
				})
				return 
			}
		}
		if err == nil {
			extension := filepath.Ext(file.Filename)
			newFileName := uuid.New().String() + extension
			newFilepath="C:/Users/kamal/Documents/SE project/Gator-X/backend/webapp/images/reviewimages/" + newFileName
			if err := c.SaveUploadedFile(file, newFilepath); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to save the file",
				})
			return
			}
		} 
		data,_:=c.GetPostForm("data")
		
		var json1 model.BaseReview
		json.Unmarshal([]byte(data), &json1)
		result := db.Model(&breview).Where("id = ? AND reviewer_id = ?", reviewid, reviewerid).Updates(&json1)

		fmt.Println(json1.Rating)
		fmt.Println(breview.PlaceID)
		if json1.Rating!=0 && breview.PlaceID!=0{
			var placereviews []model.BaseReview
			db.Find(&placereviews, "place_id = ?", breview.PlaceID)
			var avgrating =float64(0.0)
			for i := 0; i < len(placereviews); i++ {
				avgrating+=float64(placereviews[i].Rating)
			}
			fmt.Println(avgrating)
			avgrating=avgrating/float64(len(placereviews))
			avgrating=math.Round(avgrating)
			var uplace model.Places
			result1 := db.Model(&uplace).Where("id = ?", breview.PlaceID).Update("avg_rating",avgrating)

			if result1.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": result1.Error.Error()})
				return
			}
			// 
		}
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

/*
	DeletereviewView - deletes the review from the database, using the current session if user logged in if not then user should login and only logged in users have
	can delete a place. Taking the reviewID to find the review in the database. if there isn't such review then return review
	doesn't exist or any issue in finding the review return the error with a status bad request, if not then first remove the image file if exists from the server and then delete
	the record from the database, return a message saying place is deleted from the database.
*/


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
		result1:=db.Find(&breview, "id = ? AND reviewer_id = ?", reviewid, reviewerid)
		if (breview==model.BaseReview{}){
			c.JSON(http.StatusBadRequest, gin.H{"error": "Review does not exist"})
			return
		}

		if result1.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result1.Error.Error()})
			return
		}
		// fmt.Println(place)
		if breview.ReviewImage!=""{
			e := os.Remove(breview.ReviewImage)
			if e != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to delete the file",
				})
				return 
			}
		}
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


/*
	GetreviewsbyuserView - returns reviews given the user id, search the database with user id then return reviews if present else return errors.

*/
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


/*
	GetreviewsbyplaceView - returns reviews for a particular place given the place id, search the database with place id, first check the database if the place exists,
	if present check reviews of that particular place in the database and return them else return errors.

*/
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