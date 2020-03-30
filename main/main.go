package main

import (
	"NewsFeeder/db"
	"NewsFeeder/endpoints"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main()  {

	var err error
	db.DB, err = gorm.Open("postgres", "host=localhost port=5432 user=akoridze dbname=postgres password=docker sslmode=disable")

	if err != nil {
		fmt.Println("ERROR ", err)
	}
	defer db.DB.Close()
	db.DB.AutoMigrate(&db.NewsFeed{})

	r := gin.Default()

	r.GET("/newsfeeds", endpoints.GetAllNews())
	r.POST("/newsfeed", endpoints.PostNewsFeed())

	_ = r.Run()
}

