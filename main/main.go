package main

import (
	"github.com/AnaKoridze/NewsFeeder/app"
	"github.com/AnaKoridze/NewsFeeder/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	app.New().Run()

	r := gin.Default()

	r.GET("/newsfeeds", handlers.GetAllNews())
	r.POST("/newsfeed", handlers.PostNewsFeed())

	_ = r.Run()
}
