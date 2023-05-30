package main

import (
	"fmt"
	"myWebApp/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

// main is the main entry point of the app
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/contact", handlers.Contact)
	fmt.Println(fmt.Sprintf("Starting the application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
