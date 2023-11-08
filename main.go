package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// in order for a function to response to a request from a browser, it needs to handle 2 parameters: ResponseWriter and Request
// in Go it is a custom to begin comments with the name of the function
// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(3, 3)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 3 + 3 is %d", sum))
}

// addValues takes two integers and returns the sum
func addValues(x, y int) int {
	return x + y
}

// main is the main application function
func main() {
	//fmt.Println("vim-go")
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	n, err := fmt.Fprintf(w, "Hello, world")
	//	if err != nil {
	//		fmt.Println("Error occurred:", err)
	//	}
	//	fmt.Println(fmt.Sprintf("Number of bytes written was: %d", n))
	//})
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting the server at: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
