package handlers

import (
	"github.com/AnaKoridze/NewsFeeder/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostNewsFeed() gin.HandlerFunc {
	return func(context *gin.Context) {
		request := models.NewsFeed{}
		err := context.BindJSON(&request)

		if err != nil {
			context.Status(http.StatusInternalServerError)
			return
		}

		if err = models.DB.Create(&request).Error; err != nil {
			context.JSON(http.StatusBadRequest, create.CreateNewsFeedResponse{0})
		} else {
			context.JSON(http.StatusOK, create.CreateNewsFeedResponse{1})
		}
	}
}
