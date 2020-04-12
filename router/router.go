package router

import (
	"github.com/AnaKoridze/NewsFeeder/handlers"
	"github.com/AnaKoridze/NewsFeeder/services"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(services *services.Services) *gin.Engine {

	// Set the default gin router
	r := gin.Default()

	// Initialize Routes
	initializeRoutes(r, services)

	return r

}


func initializeRoutes(r *gin.Engine, services *services.Services) {

	r.GET("/newsfeeds", handlers.GetAllNews(services))
	r.POST("/newsfeed", handlers.PostNewsFeed(services))

	_ = r.Run()
}
