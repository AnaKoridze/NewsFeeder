package endpoints

import (
	"NewsFeeder/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllNews(feed *data.Repo) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, feed.GetAllNews())
	}
}