package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	// Our web app had two routes
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting the application on port %s", portNumber))

	// web app is listening on port 8080 for request
	http.ListenAndServe(portNumber, nil)
}
