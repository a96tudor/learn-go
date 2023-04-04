package render

import (
	"log"
	"net/http"
	"path"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles(path.Join("templates", tmpl))
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		log.Println(err)
	}
}
