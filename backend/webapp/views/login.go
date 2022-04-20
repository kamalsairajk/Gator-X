package views

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"

	model "webapp/model"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
)


/*
	Register View - registers a new user, First try to bind the requested json to the User Struct and in the case of of field names 
	missing or are wrong, return a bad request, strip HTML input from strings using the Strip tag policy to get the relevant field 
	details and set type to be normal, if the user exists, then return a Status Conflict, with a error message User Already exists.
	If passed, then create the user.	

*/
func RegisterView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var json model.Users
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		p := bluemonday.StripTagsPolicy()
		json.Name = p.Sanitize(json.Name)
		json.Email = p.Sanitize(json.Email)
		json.Phone = p.Sanitize(json.Phone)
		json.Password = p.Sanitize(json.Password)
		json.Type=model.NORMAL

		var user model.Users
		db.Find(&user, "email = ?", json.Email)
		if user != (model.Users{}) {
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

/*
	Login View - logs in user, create a session with default context returns a Status Bad request if not binded, take in the entered 
	email and password by Stripping the tags, search the users table with email and password, if found, set session id and return 
	a success login message else return a Status unauthorized with error message invalid username and password.
*/
func LoginView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		session := sessions.Default(c)
		var json model.Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		p := bluemonday.StripTagsPolicy()
		json.Email = p.Sanitize(json.Email)
		json.Password = p.Sanitize(json.Password)

		var user []model.Users
		db.Find(&user, "email = ? AND password = ?", json.Email, json.Password)
		if len(user) > 0 {
			session.Set("userId", user[0].ID)
			session.Save()
			c.JSON(http.StatusOK, gin.H{
				"result": "login success",
			})
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username and password",
		})
	}
	return gin.HandlerFunc(fn)

}

/*
	Logout View - remove the user from session, check login status if not logged in return status unauthorized message, else
	clear the session and save it. If these dont work return logout failed with status unauthorized message.
*/
func LogoutView(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("userId")
	if userId == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": "Login Required",
		})
		return
	}
	session.Clear()
	session.Save()
	logoutuser := session.Get("userId")
	if logoutuser == nil {
		c.JSON(http.StatusOK, gin.H{
			"result": "Logout successful",
		})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"result": "Logout Failed",
	})

}
/*
	GetUserbyID view - return user by user ID, take in the parameter userID, use this to find the user in the user table in the database.
	If found return the user with result as message in the object else if not found then return the error message. 
*/
func GetUserbyIDView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userid, _ := strconv.Atoi(c.Param("userID"))

		var user model.Users
		result := db.Find(&user, "id = ?", userid)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": user,
		})
	}
	return gin.HandlerFunc(fn)
}

/*
	DeleteUser view - deletes the user from the database, by taking in the parameter of userID, using this to find the user in the table in the
	database. If found deletes the record from the database and returns result message with record deleted and returns the http status bad request. 

*/
func DeleteUserView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userid, _ := strconv.Atoi(c.Param("userID"))

		var user model.Users
		db.Find(&user, "id = ?", userid)
		result := db.Delete(&user)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": "User deleted from database",
		})
	}
	return gin.HandlerFunc(fn)
}

/*
	Getallusers view -  return users present in the database, get the users with a model object array to store all users.
	If found return the users with result as message in the object else if not found then return the empty array message.
*/
func GetallusersView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var users []model.Users

		db.Find(&users)

		if len(users) > 0 {
			c.JSON(http.StatusOK, gin.H{
				"result": users,
			})
			return
		}
		if len(users) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "currently the database is empty",
				"result":  users,
			})
			return
		}

	}

	return gin.HandlerFunc(fn)
}
/*
	Register View - registers a admin user, First try to bind the requested json to the User Struct and in the case of of field names 
	missing or are wrong, return a bad request, strip HTML input from strings using the Strip tag policy to get the relevant field 
	details and set type to be admin, if the user exists, then return a Status Conflict, with a error message User Already exists.
	If passed, then create the user.	

*/


func RegisterAdminView(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var json model.Users
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		p := bluemonday.StripTagsPolicy()
		json.Name = p.Sanitize(json.Name)
		json.Email = p.Sanitize(json.Email)
		json.Phone = p.Sanitize(json.Phone)
		json.Password = p.Sanitize(json.Password)
		json.Type=model.ADMIN

		var user model.Users
		db.Find(&user, "email = ?", json.Email)
		if user != (model.Users{}) {
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

