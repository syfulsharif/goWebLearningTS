package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the entry point of the application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Application is running at localhost%s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
