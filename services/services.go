package services

import "github.com/AnaKoridze/NewsFeeder/config"

type Services struct {
	NewsFeedService NewsFeedRepository
}

func InitServices(config config.AppConfig) *Services {

	newsFeedService, err := initNewsFeedService(config)
	if err != nil {
		println("newsFeedService init failed ", err.Error())
	}

	return &Services{newsFeedService}
}