package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLiveness(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{"liveness": "OK"})
}

