package handlers

import (
	"helloTemplate/pkg/render"
	"net/http"
)

// Home is the homepage Handler of this web app!
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
