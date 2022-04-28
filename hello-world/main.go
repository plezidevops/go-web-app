package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))

}

// addValues adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(2.0, 0.0)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, fmt.Sprintf(" %.2f divided by %.2f is %.2f", 2.0, 0.0, f))
}

func divideValues(x, y float64) (float64, error) {

	if y <= 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}
	result := x / y

	return result, nil
}

func main() {
	// routes
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting the application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
