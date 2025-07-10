package models

import (
	"time"

	"gorm.io/gorm"
)

type Methods string

const (
	GET    Methods = "GET"
	POST   Methods = "POST"
	PUT    Methods = "PUT"
	DELETE Methods = "DELETE"
)

type TaskStatus string

const (
	TaskStatusActive    TaskStatus = "active"
	TaskStatusPaused    TaskStatus = "paused"
	TaskStatusFailed    TaskStatus = "failed"
	TaskStatusCompleted TaskStatus = "completed"
)

type Task struct {
	gorm.Model
	UserID             uint              `gorm:"not null;index" json:"user_id"`
	Name               string            `gorm:"size:100;not null" json:"name"`
	Description        string            `gorm:"size:500" json:"description"`
	ScheduleExpression string            `gorm:"size:100;not null" json:"schedule_expression"`
	Timezone           string            `gorm:"size:15;default:'UTC-6'" json:"timezone"`
	Endpoint           string            `gorm:"not null" json:"endpoint"`
	Method             Methods           `gorm:"type:enum('GET','POST','PUT','DELETE');default:'GET'" json:"method"`
	Headers            map[string]string `gorm:"serializer:json" json:"headers"`
	Body               string            `gorm:"type:text" json:"body"`
	QueryParameters    map[string]string `gorm:"serializer:json" json:"query_parameters"`
	Status             TaskStatus        `gorm:"type:enum('active','paused','failed','completed');default:'active'" json:"status"`
	Repeat             bool              `gorm:"default:false" json:"repeat"`
	LastRunAt          *time.Time        `json:"last_run_at"`
	LastRunStatus      string            `gorm:"size:20" json:"last_run_status"`
	NextRunAt          time.Time         `gorm:"index" json:"next_run_at"`
}
