package main


// imports model,views,Gin and Gorm


import (
	model "webapp/model"
	view "webapp/views"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// middleware to set headers 

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

/*
	Setup the database in order to connect the database with default configuration, if failed return a panic with message failed to connect database
	and setup all tables

*/
func db_setup(dbname string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.BaseReview{})
	db.AutoMigrate(&model.Places{})
	db.AutoMigrate(&model.Users{})
	return db
}
/*
	Setup the webserver with default configuration and pass through the CORS middleware to check every request. Logger added to the middleware and recovery to use the 
	default recovery mechanism to counter any unexpected crashes in web server. And endpoints that server as an API for the frontend.
*/

func backendserver_setup(db *gorm.DB, cookiestorename string, sessionname string) *gin.Engine {
	server := gin.Default()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.Use(CORSMiddleware())

	cookiesstore := cookie.NewStore([]byte(cookiestorename))
	cookiesstore.Options(sessions.Options{MaxAge: 60 * 60 * 24})
	server.Use(sessions.Sessions(sessionname, cookiesstore))

	server.POST("/register", view.RegisterView(db))
	server.POST("/login", view.LoginView(db))
	server.POST("/logout", view.LogoutView)
	server.GET("/users/:userID", view.GetUserbyIDView(db))
	server.DELETE("/deleteuser/:userID", view.DeleteUserView(db))
	server.GET("/getallusers", view.GetallusersView(db))	
	server.POST("/registeradmin",view.RegisterAdminView(db))  
	
	server.GET("/getallplaces", view.GetallplacesView(db))
	server.POST("/postplace", view.PostplaceView(db)) 
	server.GET("/getplace/:placeID", view.GetPlacebyIDView(db)) 
	server.PATCH("/editplace/:placeID", view.EditplaceView(db))  
	server.DELETE("/deleteplace/:placeID", view.DeleteplaceView(db)) 

	server.GET("/getallreviews", view.GetallreviewsView(db))
	server.POST("/postreview", view.PostreviewView(db))
	server.PATCH("/editreview/:reviewID", view.EditreviewView(db))
	server.DELETE("/deletereview/:reviewID", view.DeletereviewView(db))	
	server.GET("/getuserreviews", view.GetreviewsbyuserView(db))  
	server.GET("/getplacereviews/:placeID", view.GetreviewsbyplaceView(db)) 
	
	return server
}
// function that starts up the server and the database.
func main() {
	db := db_setup("main.db")
	server := backendserver_setup(db, "secret", "newsession")
	server.Run()

}
