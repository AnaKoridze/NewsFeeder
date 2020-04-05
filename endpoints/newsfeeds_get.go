package endpoints

import (
	"github.com/AnaKoridze/NewsFeeder/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllNews() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, db.DB.Find(&db.NewsFeeds))
	}
}