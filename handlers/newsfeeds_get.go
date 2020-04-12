package handlers

import (
	"github.com/AnaKoridze/NewsFeeder/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllNews(services *services.Services) gin.HandlerFunc {
	return func(context *gin.Context) {
		response, err := services.NewsFeedService.GetAllFeeds()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"error":  err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, response)
	}
}
