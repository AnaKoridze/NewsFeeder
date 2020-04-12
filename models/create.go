package models

type CreateNewsFeedRequest struct {
	Title  string `json:"title" binding:"required"`
	Post   string `json:"post" binding:"required"`
	Author string `json:"author,omitempty"`
}

type CreateNewsFeedResponse struct {
	ID uint `json:"id"`
}
