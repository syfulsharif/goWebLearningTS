package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)

	if err != nil {
		fmt.Fprintf(w, "Can not divide by 0.0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("<h2>%f divided by %f is %f</h2>", 100.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("can not divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// main is the entry point of the application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Application is running at localhost%s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
