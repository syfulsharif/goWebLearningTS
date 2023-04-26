package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

// Home is the homepage Handler of this web app!
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about.page.tmpl")
}

// main is the entry point of the application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Application is running at localhost%s", portNumber))
	http.ListenAndServe(portNumber, nil)
}

func renderTemplates(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
