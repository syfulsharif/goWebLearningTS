package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the homepage handler
func Home(w http.ResponseWriter, r *http.Request) {

}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {

}

// main is the main entry point of the app
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("Starting the application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
