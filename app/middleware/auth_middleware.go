package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/LuisAGP/cronjobs/app/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		token, err := c.Cookie("access_token")

		if err != nil {
			session.Set("error", "No hay una sesión activa")
			session.Save()
			c.Redirect(http.StatusFound, "/login")
			return
		}

		claims, err := auth.ValidateToken(token)

		if err != nil {
			session.Set("error", "No hay una sesión activa")
			session.Save()
			c.Redirect(http.StatusFound, "/login")
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
		return
	}
}

func AuthMiddlewareAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Autorización requerida"})
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			return
		}

		claims, err := auth.ValidateToken(tokenParts[1])
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		// Guardar información del usuario en el contexto
		c.Set("userId", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()
	}
}
