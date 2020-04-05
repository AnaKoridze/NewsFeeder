package main

import (
	"fmt"
	"github.com/AnaKoridze/NewsFeeder/config"
	"github.com/AnaKoridze/NewsFeeder/db"
	"github.com/AnaKoridze/NewsFeeder/endpoints"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main()  {

	conf := config.LoadConfig()

	connection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.PostgresqlHost,
		conf.PostgresqlPort,
		conf.PostgresqlUser,
		conf.PostgresqlPassword,
		conf.PostgresqlDbName)

	// open connection
	var err error
	db.DB, err = gorm.Open("postgres", connection)

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

