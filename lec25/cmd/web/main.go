package main

import (
	"fmt"
	"helloTemplate/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

// main is the entry point of the application
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Application is running at localhost%s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
