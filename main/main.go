package main

import (
	"NewsFeeder/data"
	"NewsFeeder/endpoints"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()

	feed := data.NewRepo()

	r.GET("/newsfeeds", endpoints.GetAllNews(feed))
	r.POST("/newsfeed", endpoints.PostNewsFeed(feed))

	_ = r.Run()
}

