package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

var reviews []model.BaseReview
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
			Type: model.NORMAL,
		},
		{
			Name:     "testuser2",
			Password: "Testuser2@456",
			Email:    "testuser2@gmail.com",
			Phone:    "+1 344 122 8777",
			Type: model.NORMAL,
		},
		{
			Name:     "testuser3",
			Password: "Testuser3@789",
			Email:    "testuser3@gmail.com",
			Phone:    "+1 222 333 4477",
			Type: model.NORMAL,
		},
		{
			Name:     "adminuser1",
			Password: "Adminuser1@334",
			Email:    "adminuser1@gmail.com",
			Phone:    "+1 332 444 6699",
			Type: model.ADMIN,
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

	reviews = []model.BaseReview{
		{

			ReviewTitle: "Good sandwiches",
			Review:      "The sandwiches are really good.",
			Rating:      3,
			PlaceID:     1,
			ReviewerID:  1,
		},
		{
			ReviewTitle: "Decent subs",
			Review:      "The subs here are not so good when compared to your usual subway.",
			Rating:      2,
			PlaceID:     2,
			ReviewerID:  2,
		},
		{
			ReviewTitle: "bad shakes",
			Review:      "The shakes here are not so great.",
			Rating:      2,
			PlaceID:     3,
			ReviewerID:  3,
		},
	}
	db.Create(&reviews)

}

//get all users pass case 
func testcase1(t *testing.T, router *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getallusers", nil)
	router.ServeHTTP(w, req)
	var a string = `{"result":`
	assert.Equal(t, 200, w.Code)
	b, _ := json.Marshal(users)
	assert.Equal(t, a+string(b)+"}", w.Body.String())
}

//get all places pass case
func testcase2(t *testing.T, router *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getallplaces", nil)
	router.ServeHTTP(w, req)
	var a string = `{"result":`
	assert.Equal(t, 200, w.Code)
	b, _ := json.Marshal(places)
	assert.Equal(t, a+string(b)+"}", w.Body.String())
}

// register user pass case
func testcase3(t *testing.T, router *gin.Engine) {
	w := httptest.NewRecorder()
	var jsonData = []byte(`{
		"name":"testuser4",
		"email":"terstuser4@gmail.com",
		"password":"Testuser4@345",
		"phone":"+1 345 678 9901"
	}`)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)
	// var a string = `{"result":`
	assert.Equal(t, 200, w.Code)
	expoutput := `{"result":"User created in database"}`
	assert.Equal(t, expoutput, w.Body.String())
}

// register user - invalid input - fail case
func testcase4(t *testing.T, router *gin.Engine) {
	w := httptest.NewRecorder()
	var jsonData = []byte(`{
		"name":"testuser4",
		"email":"testuser4@gmail.com",
		"password":"Testuser4@345",
		"phone":""
	}`)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

}

//post user review - unauthorized - fail case
func testcase5(t *testing.T, router *gin.Engine) {
	w := httptest.NewRecorder()
	var jsonData = []byte(`{
		"reviewtitle":"good sandwiches",
		"review":" The food here is good especially sandwiches",
		"rating":3,
		"placeid":1
	}`)
	req, _ := http.NewRequest("POST", "/postreview", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)
	// var a string = `{"result":`
	assert.Equal(t, 400, w.Code)
	expoutput := `{"error":"user not logged in"}`
	assert.Equal(t, expoutput, w.Body.String())
}

//edit user review - pass case
func testcase6(t *testing.T, router *gin.Engine) {
	w := httptest.NewRecorder()
	var jsonData1 = []byte(`{
		"password": "Testuser3@789",
		"email":    "testuser3@gmail.com"

	}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData1))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("credentials", "include")
	router.ServeHTTP(w, req)
	cookieValue := w.Result().Header.Get("Set-Cookie")
	if w.Code == 200 {
		var jsonData2 = []byte(`{
			"reviewtitle":"bad sandwiches",
			"review":" The food here is bad especially sandwiches",
			"rating":1,
			"placeid":1
		}`)
		w.Flush()
		req, _ := http.NewRequest("PATCH", "/editreview/1", bytes.NewBuffer(jsonData2))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("credentials", "include")
		req.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		// expoutput := `"result":"Review edited in database"`
		// assert.Equal(t, expoutput, w.Body.String())
	}

}

//delete user - pass case
func testcase7(t *testing.T, router *gin.Engine) {
	w := httptest.NewRecorder()
	var jsonData1 = []byte(`{
		"password": "Testuser3@789",
		"email":    "testuser3@gmail.com"

	}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData1))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("credentials", "include")
	router.ServeHTTP(w, req)
	cookieValue := w.Result().Header.Get("Set-Cookie")
	if w.Code == 200 {
		w.Flush()
		req, _ := http.NewRequest("DELETE", "/users/2", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("credentials", "include")
		req.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		// expoutput := `"result":"Review edited in database"`
		// assert.Equal(t, expoutput, w.Body.String())
	}
}

// create place - pass case
func testcase8(t *testing.T, router *gin.Engine) {
	w:= httptest.NewRecorder()
	var jsonData1 = []byte(`{
		"password": "Adminuser1@334",
		"email":    "adminuser1@gmail.com"

	}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData1))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("credentials", "include")
	router.ServeHTTP(w, req)
	cookieValue := w.Result().Header.Get("Set-Cookie")
	if w.Code == 200 {
		nr:=httptest.NewRecorder()
		var jsonData = []byte(`{
				"placename":"Shake Smart",
				"location":"Reitz Union, UF",
				"type":"Beverage",
				"avgrating":3
		}`)
		req1, _ := http.NewRequest("POST", "/postplace", bytes.NewBuffer(jsonData))
		req.Header.Set("credentials", "include")
		req1.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr,req1)
		// var a string = `{"result":`
		assert.Equal(t, 200, nr.Code)
		expoutput := `{"result":"Place created in database"}`
		assert.Equal(t, expoutput, nr.Body.String())

	}
	
	

}

//user login - pass case
func testcase9(t *testing.T, router *gin.Engine) {
	w := httptest.NewRecorder()
	var jsonData1 = []byte(`{
		"password": "Testuser1@123",
		"email":    "testuser1@gmail.com"

	}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData1))
	router.ServeHTTP(w, req)
	// var a string = `{"result":`
	assert.Equal(t, 200, w.Code)
	expoutput := `{"result":"login success"}`
	assert.Equal(t, expoutput, w.Body.String())
}

//user login - wrong login information - fail case 
func testcase10(t *testing.T, router *gin.Engine) {
	w := httptest.NewRecorder()
	var jsonData1 = []byte(`{
		"password": "Testuser2@456",
		"email":    "wrond@gmail.com"

	}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData1))
	router.ServeHTTP(w, req)
	// var a string = `{"result":`
	assert.Equal(t, 401, w.Code)
	// expoutput := `{"result":"login success"}`
	fmt.Println(w.Body.String())
	// assert.Equal(t, expoutput, w.Body.String())
}

//get place by id -- pass case
func testcase11(t *testing.T, router *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getplace/1", nil)
	router.ServeHTTP(w, req)
	var a string = `{"result":`
	assert.Equal(t, 200, w.Code)
	b, _ := json.Marshal(places[0])
	assert.Equal(t, a+string(b)+"}", w.Body.String())
}

//delete place -- pass case
func testcase12(t *testing.T, router *gin.Engine) {
	
	w:= httptest.NewRecorder()
	var jsonData1 = []byte(`{
		"password": "Adminuser1@334",
		"email":    "adminuser1@gmail.com"

	}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData1))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("credentials", "include")
	router.ServeHTTP(w, req)
	cookieValue := w.Result().Header.Get("Set-Cookie")
	if w.Code == 200 {
		nr := httptest.NewRecorder()
		req1, _ := http.NewRequest("DELETE", "/deleteplace/2", nil)
		req.Header.Set("credentials", "include")
		req1.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr, req1)
		assert.Equal(t, 200, nr.Code)
		expoutput := `{"result":"Place deleted from database"}`
		fmt.Println(nr.Body.String())
		assert.Equal(t, expoutput, nr.Body.String())

	}

}

//edit place --place not exists - fail case 
func testcase13(t *testing.T, router *gin.Engine) {

	w:= httptest.NewRecorder()
	var jsonData1 = []byte(`{
		"password": "Adminuser1@334",
		"email":    "adminuser1@gmail.com"

	}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData1))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("credentials", "include")
	router.ServeHTTP(w, req)
	cookieValue := w.Result().Header.Get("Set-Cookie")
	if w.Code == 200 {
		nr := httptest.NewRecorder()
		req1, _ := http.NewRequest("EDIT", "/editplace/4", nil)
		req.Header.Set("credentials", "include")
		req1.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr, req1)
		assert.Equal(t, 404, nr.Code)
		// expoutput := `{"result":"Place deleted from database"}`
		// fmt.Println(nr.Body.String())
		// assert.Equal(t, expoutput, nr.Body.String())

	}
}

//getuserreviews -- pass case
func testcase14(t *testing.T, router *gin.Engine) {
	w:= httptest.NewRecorder()
	var jsonData1 = []byte(`{
		"password": "Testuser3@789",
		"email":    "testuser3@gmail.com"

	}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData1))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("credentials", "include")
	router.ServeHTTP(w, req)
	cookieValue := w.Result().Header.Get("Set-Cookie")
	if w.Code == 200 {
		nr:=httptest.NewRecorder()
		req1, _ := http.NewRequest("GET", "/getuserreviews", nil)
		// req1.Header.Set("Content-Type", "application/json")
		req.Header.Set("credentials", "include")
		req1.Header.Set("Cookie", cookieValue)
		router.ServeHTTP(nr,req1)
		assert.Equal(t, 200, nr.Code)
		var a string = `{"result":[`
		b, _ := json.Marshal(reviews[2])
		assert.Equal(t, a+string(b)+"]}", nr.Body.String())

	}
}
//getplacereviews - pass case 
func testcase15(t *testing.T, router *gin.Engine) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getplacereviews/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var a string = `{"result":[`
	b, _ := json.Marshal(reviews[0])
	assert.Equal(t, a+string(b)+"]}", w.Body.String())
}

//registeradmin - pass case
func testcase16(t *testing.T, router *gin.Engine){
	w := httptest.NewRecorder()
	var jsonData = []byte(`{
		"name":"adminuser2",
		"email":"adminuser2@gmail.com",
		"password":"Adminuser2@345",
		"phone":"1234567890"
	}`)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}



func TestAllcases(t *testing.T) {

	db := testdb_setup("test.db")

	initData(db)

	router := backendserver_setup(db, "teststore", "testsession")

	testcase1(t, router)
	testcase2(t, router)
	testcase3(t, router)
	testcase4(t, router)
	testcase5(t, router)
	testcase6(t, router)
	testcase7(t, router)
	testcase8(t, router)
	testcase9(t, router)
	testcase10(t, router)
	testcase11(t,router)
	testcase12(t,router)
	testcase13(t,router)
	testcase14(t,router)
	testcase15(t,router)
	testcase16(t,router)

}
