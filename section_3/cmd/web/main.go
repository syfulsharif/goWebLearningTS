package main

import (
	"fmt"
	"log"
	"myWebApp/pkg/config"
	"myWebApp/pkg/handlers"
	"myWebApp/pkg/render"
	"net/http"
)

const portNumber = ":8080"

// main is the main entry point of the app
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not create templae cache")
	}
	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/contact", handlers.Contact)
	fmt.Println(fmt.Sprintf("Starting the application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
