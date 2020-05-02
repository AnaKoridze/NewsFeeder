package services

import "github.com/AnaKoridze/NewsFeeder/models"

type NewsFeedRepository interface {
	GetAllFeeds() (*models.GetNewsFeedsResponse, error)
	CreateFeed(request models.CreateNewsFeedRequest) (*models.CreateNewsFeedResponse, error)
}
