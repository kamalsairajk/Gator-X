package views

import (
	// "fmt"
	"net/http"
	"strconv"
	"path/filepath"
	model "webapp/model"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/gorm"
	"github.com/gin-contrib/sessions"
	"encoding/json"
	"os"


)


/*
	GetallplacesView - return all places, using an array of places model to store the results, find the database, and if the length of places greater than
	0 return the results and if length of places is 0 then return a message that the table is currently empty.
*/
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

/*
	PostplaceView - create a place with given details, using the current session if user logged in if not then user should login and only admin users have
	this feature to create a place, all place related details are added into a json object under the form field data and image related to the place is added
	into the file field of the form data, after performing some data processing the data is stored into the database and if the place already exists return a
	message of place already exists.
*/
func PostplaceView(db *gorm.DB) gin.HandlerFunc {
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

		file, err := c.FormFile("file")
		newFilepath:=""
		if err == nil {
			extension := filepath.Ext(file.Filename)
			newFileName := uuid.New().String() + extension
			newFilepath="C:/Users/kamal/Documents/SE project/Gator-X/backend/webapp/images/placeimages/" + newFileName
			if err := c.SaveUploadedFile(file, newFilepath); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to save the file",
				})
			return
			}
		} 

		data,_:=c.GetPostForm("data")

		var json1 model.Places
		json.Unmarshal([]byte(data), &json1)
		// if err:=json.Unmarshal([]byte(data), &json1);err!=nil{
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//  	return
		// }

		p := bluemonday.StripTagsPolicy()
		json1.Placename = p.Sanitize(json1.Placename)
		json1.Location = p.Sanitize(json1.Location)
		json1.Type = p.Sanitize(json1.Type)
		json1.AvgRating, _ = strconv.Atoi(p.Sanitize(strconv.Itoa(json1.AvgRating)))
		json1.PlaceImage=newFilepath

		var place model.Places
		db.Find(&place, "placename = ? AND location = ? AND type = ?", json1.Placename, json1.Location, json1.Type)
		if place != (model.Places{}) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Place in that location Already Exists!"})
			return
		}

		result := db.Create(&json1)

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

/*
	GetPlacebyIDView - return a place when ID of the place is provided, placeID provided as a parameter, having a model object to store the results, when
	queried through the database with the given id if not found return the error message else if found return the place.
*/

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


/*
	DeleteplaceView - deletes the place from the database, using the current session if user logged in if not then user should login and only admin users have
	this feature to delete a place. Taking the placeID of a particular place to find the place in the database. if there isn't such place then return place
	doesn't exist or any issue in finding the place return the error with a status bad request, if not then first remove the file from the server and then delete
	the record from the database, return a message saying place is deleted from the database.
*/
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
		// fmt.Println(placeid)
		var place model.Places
		result1:=db.Find(&place, "id = ?", placeid)
		
		if (place==model.Places{}){
			c.JSON(http.StatusBadRequest, gin.H{"error": "Place id does not exist"})
			return
		}

		if result1.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result1.Error.Error()})
			return
		}
		// fmt.Println(place)
		if place.PlaceImage!=""{
			e := os.Remove(place.PlaceImage)
			if e != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to delete the file",
				})
				return 
			}
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

/*
	EditplaceView - edit a place details with given constraints or fields, using the current session if user logged in if not then user should login and only admin users have
	this feature to edit a place. Taking the place id as a parameter to edit and find the place in the database, if the place is not present then return a status bad request
	else if finding the place returns an error then return the error message, if found then update the record in the database with the given details and if file is present 
	delete existing place in the server and save the current file in the server and change this in the database along with other details. Finally return an message that says
	the place is edited.
*/
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
		// fmt.Println(placeid)
		var place1 model.Places
		result1 := db.Find(&place1, "id = ?", placeid)

		if (place1==model.Places{}){
			c.JSON(http.StatusBadRequest, gin.H{"error": "Place id does not exist"})
			return
		}

		if result1.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result1.Error.Error()})
			return
		}
		// fmt.Print(result1)
		
		file, err := c.FormFile("file")
		newFilepath:=""
		if place1.PlaceImage!="" && err==nil{
			// fmt.Println(place1.PlaceImage)
			e := os.Remove(place1.PlaceImage)
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
			newFilepath="C:/Users/kamal/Documents/SE project/Gator-X/backend/webapp/images/placeimages/" + newFileName
			if err := c.SaveUploadedFile(file, newFilepath); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to save the file",
				})
			return
			}
		} 
		data,_:=c.GetPostForm("data")

		var place model.Places
		json.Unmarshal([]byte(data), &place)

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




