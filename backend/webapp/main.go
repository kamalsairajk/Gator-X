package main

import (
	model "webapp/model"
	view "webapp/views"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.BaseReview{})
	db.AutoMigrate(&model.Places{})
	db.AutoMigrate(&model.Users{})

	server := gin.Default()

	server.Use(gin.Logger())

	server.Use(gin.Recovery())
	cookiesstore := cookie.NewStore([]byte("secret"))
	cookiesstore.Options(sessions.Options{MaxAge: 60 * 60 * 24})
	server.Use(sessions.Sessions("newsession", cookiesstore))

	server.GET("/getallplaces", view.GetallplacesView(db))
	server.POST("/postplace", view.PostplaceView(db))
	server.GET("/getallreviews", view.GetallreviewsView(db))
	server.POST("/postreview", view.PostreviewView(db))
	server.PATCH("/editreview/:reviewID", view.EditreviewView(db))
	server.DELETE("/delete/:reviewID", view.DeletereviewView(db))
	server.POST("/register", view.RegisterView(db))
	server.POST("/login", view.LoginView(db))
	server.POST("/logout", view.LogoutView)
	server.GET("/users/:userID", view.GetUserbyIDView(db))
	server.DELETE("/users/:userID", view.DeleteUserView(db))

	server.Run()

}
