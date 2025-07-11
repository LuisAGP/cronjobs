package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/LuisAGP/cronjobs/app/handlers"
	"github.com/LuisAGP/cronjobs/app/middleware"
	"github.com/LuisAGP/cronjobs/app/services"
	"github.com/LuisAGP/cronjobs/migrations"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	ginMode := os.Getenv("GIN_MODE")

	gin.SetMode(ginMode)
	r := gin.Default()

	// Middleware de sesión
	secret := os.Getenv("JWT_SECRET")
	store := cookie.NewStore([]byte(secret))
	r.Use(sessions.Sessions("sessions", store))

	// Se aplican migraciones
	args := os.Args
	if len(args) > 1 {
		if args[1] == "--migrate" {
			migrations.ApplyMigrations()
			fmt.Fprintln(os.Stdout, "Finalizado.")
			os.Exit(0)
		}
	}

	// Cargar plantillas
	services.SetHTMLTemplates(r)

	// Configurar archivos estáticos
	r.Static("/static", "./static")

	// Servir favicon.ico
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	// Middleware para inyectar la base de datos
	db := services.DB()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	/*************************************** RUTAS PÁGINA ****************************************/

	// Rutas públicas
	r.GET("/", handlers.HomeView)
	r.GET("/login", handlers.LoginView)
	r.POST("/login", handlers.Login)

	// Grupo de rutas protegidas
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/logout", handlers.Logout)
		protected.GET("/dashboard", handlers.DashboardView)
		protected.GET("/tasks", handlers.TasksView)
		protected.GET("/logs", handlers.LogsView)
		protected.GET("/docs", handlers.DocsView)

		protected.POST("/tasks", handlers.SaveTask)
	}

	/*********************************************************************************************/

	/***************************************** RUTAS API *****************************************/

	// Rutas públicas
	r.POST("/api/register", handlers.Register)
	r.POST("/api/login", handlers.Login)

	// Grupo de rutas protegidas
	protectedAPI := r.Group("/api")
	protectedAPI.Use(middleware.AuthMiddlewareAPI())
	{
		protectedAPI.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusAccepted, gin.H{"message": "pong"})
		})

		protectedAPI.POST("/tasks", handlers.SaveTask)
	}

	/*********************************************************************************************/

	host := os.Getenv("APP_HOST")
	r.Run(host)
}
