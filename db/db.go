package db

import "github.com/jinzhu/gorm"

var DB *gorm.DB

type NewsFeed struct {
	gorm.Model
	Title string `json:"title"`
	Post string `json:"post"`
}

var NewsFeeds []NewsFeed

func (n *NewsFeed) TableName() string {
	return "Newsfeed"
}