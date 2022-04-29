package main

import (
	"fmt"
	"github.com/plezidevops/go-course/pkg/handlers"
	"net/http"

	"github.com/plezidevops/go-course/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	// Our web app had two routes
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting the application on port %s", portNumber))

	// web app is listening on port 8080 for request
	http.ListenAndServe(portNumber, nil)
}
