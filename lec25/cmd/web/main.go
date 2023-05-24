package main

import (
	"fmt"
	"helloTemplate/pkg/config"
	"helloTemplate/pkg/handlers"
	"helloTemplate/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the entry point of the application
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
