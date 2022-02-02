package views

import (
	"fmt"
	"net/http"
	model "webapp/model"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
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
	}

	return gin.HandlerFunc(fn)
}
func PostplaceView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var json model.Places
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(json)

		p := bluemonday.StripTagsPolicy()
		placename := p.Sanitize(json.Placename)
		location := p.Sanitize(json.Location)
		type1 := p.Sanitize(json.Type)

		var place model.Places
		db.Find(&place, "placename = ? AND location = ? AND type = ?", placename, location, type1)
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
