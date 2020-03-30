package endpoints

import (
	"NewsFeeder/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostNewsFeed() gin.HandlerFunc {
	return func(context *gin.Context) {
		request := db.NewsFeed{}
		err := context.BindJSON(&request)

		if err != nil {
			context.Status(http.StatusInternalServerError)
			return
		}

		if err = db.DB.Create(&request).Error; err != nil {
			fmt.Println("ERROR ", err)
		}

		context.JSON(http.StatusOK, ResponseOk{"item bien ajouté"})
	}
}