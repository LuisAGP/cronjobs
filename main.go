package main

import (
	"net/http"
	"os"

	"github.com/LuisAGP/cronjobs/app/handlers"
	"github.com/LuisAGP/cronjobs/app/middleware"
	"github.com/LuisAGP/cronjobs/app/services"
	"github.com/LuisAGP/cronjobs/migrations"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	ginMode := os.Getenv("GIN_MODE")

	gin.SetMode(ginMode)
	r := gin.Default()

	// Se aplican migraciones
	args := os.Args
	if len(args) > 1 {
		if args[1] == "--migrate" {
			migrations.ApplyMigrations()
		}
	}

	// Cargar plantillas
	services.SetHTMLTemplates(r)

	// Configurar archivos estáticos
	r.Static("/static", "./static")

	// Middleware para inyectar la base de datos
	db := services.DB()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Rutas públicas
	r.GET("/login", handlers.LoginView)
	r.POST("/api/register", handlers.Register)
	r.POST("/api/login", handlers.Login)

	// Grupo de rutas protegidas
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusAccepted, gin.H{"message": "pong"})
		})
		// protected.GET("/tasks", handlers.GetTasks)
		// protected.POST("/tasks", handlers.CreateTask)
	}

	r.Run(":8080")
}
