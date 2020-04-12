package handlers

import (
	"github.com/AnaKoridze/NewsFeeder/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllNews() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, models.DB.Find(&models.NewsFeeds))
	}
}
