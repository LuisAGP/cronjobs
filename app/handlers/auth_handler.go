package handlers

import (
	"net/http"

	"github.com/LuisAGP/cronjobs/app/auth"
	"github.com/LuisAGP/cronjobs/app/inputs"
	"github.com/LuisAGP/cronjobs/app/models"
	"github.com/LuisAGP/cronjobs/app/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginView(c *gin.Context) {
	services.View(c, "login", gin.H{
		"Title": "login",
	})
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input inputs.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	if !auth.CheckPassword(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input inputs.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar si el usuario ya existe
	var existingUser models.User
	if db.Where("email = ?", input.Email).First(&existingUser).Error == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "El correo ya está registrado"})
		return
	}

	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error procesando contraseña"})
		return
	}

	newUser := models.User{
		Email:    input.Email,
		Password: hashedPassword,
	}

	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando usuario"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario registrado exitosamente"})
}
