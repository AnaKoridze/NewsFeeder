package main

import (
	"NewsFeeder/data"
	"NewsFeeder/endpoints"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()

	feed := data.NewRepo()
	feed.Add(data.Newsfeed{"bingo", "bango"})

	r.GET("/newsfeeds", endpoints.GetAllNews(feed))

	_ = r.Run()
}

