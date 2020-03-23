package endpoints

import (
	"NewsFeeder/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

type postNewsFeedRequest struct {
	Title string `json:"title"`
	Post string `json:"post"`
}

func PostNewsFeed(feed *data.Repo) gin.HandlerFunc {
	return func(context *gin.Context) {
		request := data.Newsfeed{}
		err := context.Bind(&request)

		if (err != nil) {
			context.Status(http.StatusNotAcceptable)
			return
		}

		feed.Add(request)

		context.Status(http.StatusOK)
	}
}