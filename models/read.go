package models

type GetNewsFeedsResponse struct {
	Total int         `json:"total"`
	Data  *[]NewsFeed `json:"data"`
}
