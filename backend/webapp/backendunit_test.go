package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	model "webapp/model"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var users []model.Users

// var reviews []model.BaseReview
var places []model.Places

func testdb_setup(dbName string) *gorm.DB {

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	db.Migrator().DropTable(&model.Users{})
	db.Migrator().DropTable(&model.BaseReview{})
	db.Migrator().DropTable(&model.Places{})

	db.AutoMigrate(&model.Users{}, &model.Places{}, &model.BaseReview{})

	return db
}

func initData(db *gorm.DB) {
	users = []model.Users{
		{
			Name:     "testuser1",
			Password: "Testuser1@123",
			Email:    "testuser1@gmail.com",
			Phone:    "+1 122 455 7990",
		},
		{
			Name:     "testuser2",
			Password: "Testuser2@456",
			Email:    "testuser2@gmail.com",
			Phone:    "+1 344 122 8777",
		},
		{
			Name:     "testuser3",
			Password: "Testuser3@789",
			Email:    "testuser3@gmail.com",
			Phone:    "+1 222 333 4477",
		},
	}
	db.Create(&users)

	places = []model.Places{
		{
			Placename: "Chikfila",
			Location:  "The Hub, UF",
			Type:      "Food",
			AvgRating: 3,
		},
		{
			Placename: "Subway",
			Location:  "Reitz Union, UF",
			Type:      "Food",
			AvgRating: 3,
		},
		{
			Placename: "Starbucks",
			Location:  "The Hub, UF",
			Type:      "Beverages",
			AvgRating: 3,
		},
	}
	db.Create(&places)

}

//get all users
func testcase1(t *testing.T, router *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getallusers", nil)
	router.ServeHTTP(w, req)
	var a string = `{"result":`
	assert.Equal(t, 200, w.Code)
	b, _ := json.Marshal(users)
	assert.Equal(t, a+string(b)+"}", w.Body.String())
}

//get all places
func testcase2(t *testing.T, router *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getallplaces", nil)
	router.ServeHTTP(w, req)
	var a string = `{"result":`
	assert.Equal(t, 200, w.Code)
	b, _ := json.Marshal(places)
	assert.Equal(t, a+string(b)+"}", w.Body.String())
}
func TestAllcases(t *testing.T) {

	db := testdb_setup("test.db")

	initData(db)

	router := backendserver_setup(db, "teststore", "testsession")

	testcase1(t, router)
	testcase2(t, router)
}
