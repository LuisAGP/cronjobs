package main

import (
	"github.com/LuisAGP/cronjobs/app/handlers"
	"github.com/LuisAGP/cronjobs/app/middleware"
	"github.com/LuisAGP/cronjobs/app/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := services.DB()

	// Middleware para inyectar la base de datos
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Rutas públicas
	r.POST("/api/register", handlers.Register)
	r.POST("/api/login", handlers.Login)

	// Grupo de rutas protegidas
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// protected.GET("/tasks", handlers.GetTasks)
		// protected.POST("/tasks", handlers.CreateTask)
	}

	r.Run(":8080")
}
