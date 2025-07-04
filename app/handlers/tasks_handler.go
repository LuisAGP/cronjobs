package handlers

import (
	"github.com/LuisAGP/cronjobs/app/services"
	"github.com/gin-gonic/gin"
)

func TasksView(c *gin.Context) {
	services.View(c, "tasks", gin.H{
		"Title": "Tasks",
	})
}
