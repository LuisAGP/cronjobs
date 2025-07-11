package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LuisAGP/cronjobs/app/inputs"
	"github.com/LuisAGP/cronjobs/app/models"
	"github.com/LuisAGP/cronjobs/app/services"
	"github.com/LuisAGP/cronjobs/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskListPageData struct {
	Tasks      []models.Task
	Total      int64
	Page       int
	Limit      int
	TotalPages int
	Status     string
	Search     string
	SortBy     string
	Order      string
}

func TasksView(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.GetUint("userID")

	// Get query parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "1"))
	status := c.Query("status")
	search := c.Query("search")
	sortBy := c.DefaultQuery("sort_by", "next_run_at")
	order := c.DefaultQuery("order", "asc")

	// Calculate offset
	offset := (page - 1) * limit

	var userTask []models.Task
	query := db.Model(&userTask).Where("user_id = ?", userID)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if search != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Get total count for pagination
	var total int64
	query.Count(&total)

	// Apply sorting
	sortOrder := sortBy
	if order == "desc" {
		sortOrder += " DESC"
	} else {
		sortOrder += " ASC"
	}
	query = query.Order(sortOrder)

	// Get paginated tasks
	var tasks []models.Task
	query.Offset(offset).Limit(limit).Find(&tasks)

	// Calculate total pages
	totalPages := (int(total) + limit - 1) / limit

	data := TaskListPageData{
		Tasks:      tasks,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		Status:     status,
		Search:     search,
		SortBy:     sortBy,
		Order:      order,
	}

	services.View(c, "tasks", gin.H{
		"Title": "Tasks",
		"Data":  data,
	})
}

func SaveTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.GetUint("userID")

	var req inputs.SaveTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Valculate cron expression
	if req.ScheduleExpression == "" {
		// Interval validation
		interval_value, err := strconv.Atoi(req.IntervalValue)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
			return
		}

		req.ScheduleExpression, err = utils.CalculateCronExpression(interval_value, req.IntervalUnit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	}

	// Calculate next run time
	nextRun, err := utils.CalculateNextRun(req.ScheduleExpression, req.Timezone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule expression"})
		return
	}

	timeout, err := strconv.Atoi(req.Timeout)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid timeout"})
		return
	}

	RetryCount, err := strconv.Atoi(req.RetryCount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Retry Count"})
		return
	}

	RetryInterval, err := strconv.Atoi(req.RetryInterval)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Retry Interval"})
		return
	}

	// Create task in database
	task := models.Task{
		UserID:             userID,
		Name:               req.Name,
		Description:        req.Description,
		ScheduleExpression: req.ScheduleExpression,
		Timezone:           req.Timezone,
		Endpoint:           req.Endpoint,
		Method:             req.Method,
		Headers:            req.Headers,
		Body:               req.Body,
		Timeout:            timeout,
		RetryCount:         RetryCount,
		RetryInterval:      RetryInterval,
		Status:             "active",
		NextRunAt:          nextRun,
	}

	if err := db.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"task":    task,
	})

}
