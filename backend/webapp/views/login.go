package views

import (
	"net/http"
	model "webapp/model"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
)

func RegisterView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var json model.User
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		p := bluemonday.StripTagsPolicy()
		json.Name = p.Sanitize(json.Name)
		json.Email = p.Sanitize(json.Email)
		json.Phone = p.Sanitize(json.Phone)
		json.Password = p.Sanitize(json.Password)

		var user model.User
		db.Find(&user, "email = ?", json.Email)
		if user != (model.User{}) {
			c.JSON(http.StatusConflict, gin.H{"error": "User Already Exists!"})
			return
		}
		result := db.Create(&json)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": "User created in database",
		})
	}
	return gin.HandlerFunc(fn)

}

func LoginView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var json model.Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		p := bluemonday.StripTagsPolicy()
		json.Email = p.Sanitize(json.Email)
		json.Password = p.Sanitize(json.Password)

		var user []model.User
		db.Find(&user, "email = ? AND password = ?", json.Email, json.Password)
		if len(user) > 0 {
			c.JSON(http.StatusOK, gin.H{
				"result": "login success",
			})
			return
		}
		if len(user) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "currently the database is empty",
				"result":  user,
			})
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username and password",
		})
	}
	return gin.HandlerFunc(fn)

}
