package handlers

import (
	"github.com/AnaKoridze/NewsFeeder/models"
	"github.com/AnaKoridze/NewsFeeder/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func PostNewsFeed(services *services.Services) gin.HandlerFunc {
	return func(context *gin.Context) {
		request := models.CreateNewsFeedRequest{}

		// validate json
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error":  "json decoding : " + err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}

		// check empty strings
		if len(strings.TrimSpace(request.Title)) == 0 {
			request.Title = ""
		}

		if len(strings.TrimSpace(request.Post)) == 0 {
			request.Post = ""
		}

		response, err := services.NewsFeedService.CreateFeed(request)
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
