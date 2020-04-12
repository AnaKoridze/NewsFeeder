package models

import "github.com/jinzhu/gorm"

type NewsFeed struct {
	gorm.Model
	Title  string `json:"title"`
	Post   string `json:"post"`
	Author string `json:"author"`
}

var NewsFeeds []NewsFeed

func (n *NewsFeed) TableName() string {
	return "Newsfeed"
}
