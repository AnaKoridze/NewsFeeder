package app

import (
	"github.com/AnaKoridze/NewsFeeder/config"
	"github.com/AnaKoridze/NewsFeeder/services"
	"github.com/gin-gonic/gin"
)

// Global properties
type App struct {
	Config   config.AppConfig
	Router   *gin.Engine
	Services *services.Services
}

/**
* Charger Configuration
* Initialiser le Router
* Initialiser les Services
 */
func (app *App) New() {

}
