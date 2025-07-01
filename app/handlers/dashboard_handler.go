package handlers

import (
	"github.com/LuisAGP/cronjobs/app/services"
	"github.com/gin-gonic/gin"
)

func DashboardView(c *gin.Context) {
	services.View(c, "dashboard", gin.H{
		"Title": "Dashboard",
	})
}
