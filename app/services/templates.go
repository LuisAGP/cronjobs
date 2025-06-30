package services

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetHTMLTemplates(r *gin.Engine) {
	// Creamos una instancia vacía para que luego esté disponible en el FuncMap
	var tmpl *template.Template

	// Primero creamos la FuncMap con una función vacía (por ahora)
	funcMap := template.FuncMap{}

	// Creamos función para renderizar template dentro de otro
	funcMap["renderPartial"] = func(name string) template.HTML {
		if name == "" {
			return template.HTML("")
		}

		var tpl bytes.Buffer
		err := tmpl.ExecuteTemplate(&tpl, name, nil)
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
	data["Tmpl"] = template
	c.HTML(http.StatusOK, "base", data)
}
