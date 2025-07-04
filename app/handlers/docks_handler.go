package handlers

import (
	"github.com/LuisAGP/cronjobs/app/services"
	"github.com/gin-gonic/gin"
)

func DocsView(c *gin.Context) {
	services.View(c, "docs", gin.H{
		"Title": "Docs",
	})
}
