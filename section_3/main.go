package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the home")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 5)
	fmt.Fprintf(w, fmt.Sprintf("This is about page and 2 + 5 is %d", sum))
}

// Divide is a function that divides 2 float numbers and returns a float
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "can not divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}

// AddValues adds two integers and returns an integer
func AddValues(x int, y int) int {
	return x + y
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("can not divide by 0")
		return 0, err
	}
	result := x / y
	return result, nil
}

// main is the main entry point of the app
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	fmt.Println(fmt.Sprintf("Starting the application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
