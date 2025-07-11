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
	funcMap := template.FuncMap{
		"renderPartial": func(name string, data any) template.HTML {
			if name == "" {
				return template.HTML("")
			}

			var tpl bytes.Buffer
			err := tmpl.ExecuteTemplate(&tpl, name, data)
			if err != nil {
				return template.HTML("<strong>Error rendering partial: </strong>" + err.Error())
			}
			return template.HTML(tpl.String())
		},
		"default": func(value, defaultValue any) any {
			if value == nil || value == "" {
				return defaultValue
			}
			return value
		},
		"seq": func(start, end int) []int {
			var sequence []int
			for i := start; i <= end; i++ {
				sequence = append(sequence, i)
			}
			return sequence
		},
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
	data["Path"] = c.FullPath()
	c.HTML(http.StatusOK, "base", data)
}
