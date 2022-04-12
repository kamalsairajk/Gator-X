package views

import (
	"fmt"
	"net/http"
	"strconv"
	// "path/filepath"
	model "webapp/model"
	// "github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
	"github.com/gin-contrib/sessions"

)

func GetallplacesView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var places []model.Places

		db.Find(&places)

		if len(places) > 0 {
			c.JSON(http.StatusOK, gin.H{
				"result": places,
			})
			return
		}
		if len(places) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "currently the database is empty",
				"result":  places,
			})
			return
		}

	}

	return gin.HandlerFunc(fn)
}
func PostplaceView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		// file, err := c.FormFile("file")
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		// 		"message": "No file is received",
		// 	})
		// 	return
		// }

		// extension := filepath.Ext(file.Filename)
		// newFileName := uuid.New().String() + extension

		// if err := c.SaveUploadedFile(file, "C:/Users/kamal/Documents/SE project/Gator-X/backend/webapp/placeimages" + newFileName); err != nil {
		// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		// 		"message": "Unable to save the file",
		// 	})
		// 	return
		// }
		
		// data,err:=c.GetPostForm("data")

		session := sessions.Default(c)
		i := session.Get("userId")
		if i == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user not logged in"})
			return
		}
		userId := i.(uint)
		var user model.Users
		db.First(&user, "id = ?", userId)
		if user.Type != model.ADMIN{
			c.JSON(http.StatusBadRequest, gin.H{"error": "not admin user"})
			return
		}
		var json model.Places
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(json)

		p := bluemonday.StripTagsPolicy()
		json.Placename = p.Sanitize(json.Placename)
		json.Location = p.Sanitize(json.Location)
		json.Type = p.Sanitize(json.Type)
		json.AvgRating, _ = strconv.Atoi(p.Sanitize(strconv.Itoa(json.AvgRating)))

		var place model.Places
		db.Find(&place, "placename = ? AND location = ? AND type = ?", json.Placename, json.Location, json.Type)
		if place != (model.Places{}) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Place in that location Already Exists!"})
			return
		}

		result := db.Create(&json)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": "Place created in database",
		})
	}
	return gin.HandlerFunc(fn)
}

func GetPlacebyIDView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		placeid, _ := strconv.Atoi(c.Param("placeID"))

		var place model.Places
		result := db.Find(&place, "id = ?", placeid)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": place,
		})
	}
	return gin.HandlerFunc(fn)
}

func DeleteplaceView(db *gorm.DB) gin.HandlerFunc {
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
		if user.Type != model.ADMIN{
			c.JSON(http.StatusBadRequest, gin.H{"error": "not admin user"})
			return
		}

		placeid, _ := strconv.Atoi(c.Param("placeID"))

		var place model.Places
		result1:=db.Find(&place, "id = ?", placeid)
		if result1.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result1.Error.Error()})
			return
		}

		result := db.Delete(&place)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": "Place deleted from database",
		})
	}
	return gin.HandlerFunc(fn)
}

func EditplaceView(db *gorm.DB) gin.HandlerFunc {
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
		if user.Type != model.ADMIN{
			c.JSON(http.StatusBadRequest, gin.H{"error": "not admin user"})
			return
		}
		placeid, _ := strconv.Atoi(c.Param("placeID"))

		var place model.Places
		if err := c.ShouldBindJSON(&place); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var uplace model.Places
		result := db.Model(&uplace).Where("id = ?", placeid).Updates(&place)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": "Place edited in database",
		})
	}
	return gin.HandlerFunc(fn)
}




