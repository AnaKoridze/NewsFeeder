package services

import (
	"errors"
	"fmt"
	"github.com/AnaKoridze/NewsFeeder/config"
	"github.com/AnaKoridze/NewsFeeder/models"
	"github.com/jinzhu/gorm"
)

type NewsFeedService struct {
	db *gorm.DB
}

func initNewsFeedService(config config.AppConfig) (*NewsFeedService, error) {

	connection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresqlHost,
		config.PostgresqlPort,
		config.PostgresqlUser,
		config.PostgresqlPassword,
		config.PostgresqlDbName)

	// open connection
	db, err := gorm.Open("postgres", connection)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not create PostgreSQL client: %s", err.Error()))
	}
	println("connected!")

	db.AutoMigrate(&models.NewsFeed{})

	return &NewsFeedService{db}, nil
}

func (feedService NewsFeedService) GetAllFeeds() (*models.GetNewsFeedsResponse, error) {
	d := feedService.db.Find(&models.NewsFeeds)
	if d.Error != nil {
		return nil, d.Error
	}

	response := models.GetNewsFeedsResponse{}
	response.Total = len(models.NewsFeeds)
	response.Data = &models.NewsFeeds

	return &response, nil
}

func (feedService NewsFeedService) CreateFeed(request models.CreateNewsFeedRequest) (*models.CreateNewsFeedResponse, error) {
	row := new(models.NewsFeed)

	feed := models.NewsFeed{}
	feed.Title = request.Title
	feed.Post = request.Post
	feed.Author = request.Author

	d := feedService.db.Create(&feed).Scan(&row)

	if d.Error != nil {
		return nil, d.Error
	}

	return &models.CreateNewsFeedResponse{row.ID}, nil
}