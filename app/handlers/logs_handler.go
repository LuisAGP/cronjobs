package handlers

import (
	"github.com/LuisAGP/cronjobs/app/services"
	"github.com/gin-gonic/gin"
)

func LogsView(c *gin.Context) {
	services.View(c, "logs", gin.H{
		"Title": "Logs",
	})
}
