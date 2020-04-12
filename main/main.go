package main

import (
	"fmt"
	"github.com/AnaKoridze/NewsFeeder/config"
	"github.com/AnaKoridze/NewsFeeder/handlers"
	"github.com/AnaKoridze/NewsFeeder/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {

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
	models.DB, err = gorm.Open("postgres", connection)

	if err != nil {
		fmt.Println("ERROR ", err)
	}
	defer models.DB.Close()
	models.DB.AutoMigrate(&models.NewsFeed{})

	r := gin.Default()

	r.GET("/newsfeeds", handlers.GetAllNews())
	r.POST("/newsfeed", handlers.PostNewsFeed())

	_ = r.Run()
}
