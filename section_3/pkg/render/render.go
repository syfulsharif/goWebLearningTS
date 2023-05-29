package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplate is a function that renders the template file
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing the template :", err)
		return
	}
}
