package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the homepage Handler of this web app!
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>This is the homepage</h1>")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, "<h1>This is the About page</h1> <h2>and 2 + 2 is %d </h2>", sum)
}

// addValues is the sum function that supplies sum value of 02 integers
func addValues(x, y int) int {
	return x + y
}

// main is the entry point of the application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("Application is running at localhost%s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
