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


type GetNewsFeedsResponse struct {
	Total int         `json:"total"`
	Data  *[]NewsFeed `json:"data"`
}

type CreateNewsFeedRequest struct {
	Title  string `json:"title" binding:"required"`
	Post   string `json:"post" binding:"required"`
	Author string `json:"author,omitempty"`
}

type CreateNewsFeedResponse struct {
	ID uint `json:"id"`
}

