package services

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/LuisAGP/cronjobs/app/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetHTMLTemplates(r *gin.Engine) {
	// Creamos una instancia vacía para que luego esté disponible en el FuncMap
	var tmpl *template.Template

	// Primero creamos la FuncMap con una función vacía (por ahora)
	funcMap := template.FuncMap{}

	// Creamos función para renderizar template dentro de otro
	funcMap["renderPartial"] = func(name string, data any) template.HTML {
		if name == "" {
			return template.HTML("")
		}

		var tpl bytes.Buffer
		err := tmpl.ExecuteTemplate(&tpl, name, data)
		if err != nil {
			return template.HTML("<strong>Error rendering partial</strong>")
		}
		return template.HTML(tpl.String())
	}

	// Cargamos las plantillas con el FuncMap
	tmpl = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.html"))

	r.SetHTMLTemplate(tmpl)
}

func View(c *gin.Context, template string, data gin.H) {
	// Extraemos usuario si existe
	db := c.MustGet("db").(*gorm.DB)
	userID, exists := c.Get("userID")

	if exists {
		var user models.User
		err := db.Where("id = ?", userID).First(&user).Error

		if err == nil {
			data["User"] = user
		}
	}

	data["Tmpl"] = template
	c.HTML(http.StatusOK, "base", data)
}
