package handlers

import (
	"github.com/LuisAGP/cronjobs/app/models"
	"github.com/LuisAGP/cronjobs/app/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TasksView(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	user_id := c.GetUint("userID")

	var userTask []models.Task
	db.Where("user_id = ?", user_id).Find(&userTask)

	services.View(c, "tasks", gin.H{
		"Title":     "Tasks",
		"UserTasks": userTask,
	})
}
