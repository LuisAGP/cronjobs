package handlers

import (
	"github.com/LuisAGP/cronjobs/app/services"
	"github.com/gin-gonic/gin"
)

func HomeView(c *gin.Context) {
	services.View(c, "home", gin.H{
		"Title": "Home",
	})
}
