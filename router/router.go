package router

import (
	"github.com/AnaKoridze/NewsFeeder/services"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(services *services.Services) *gin.Engine {

	// Set the default gin router
	r := gin.New()
	r.Use(gin.Recovery())

	// Initialize Routes
	initializeRoutes(r, services)

	return r

}


func initializeRoutes(r *gin.Engine, services *services.Services) {


}
