package app

import (
	"github.com/AnaKoridze/NewsFeeder/config"
	"github.com/AnaKoridze/NewsFeeder/router"
	"github.com/AnaKoridze/NewsFeeder/services"
	"github.com/gin-gonic/gin"
)

// Global properties
type App struct {
	Config   config.AppConfig
	Router   *gin.Engine
	Services *services.Services
}

func New() *App {
	app := &App{}
	app.setup()
	return app
}

/**
* Charger Configuration
* Initialiser le Router
* Initialiser les Services
 */
func (app *App) setup() {

	// Load configuration
	c := config.LoadConfig()

	// Initialize Services
	s := services.InitServices(c)

	// Initialize Router
	r := router.InitializeRouter(s)

	app.Config = c
	app.Router = r
	app.Services = s

}

func (app *App) Run() {

}
