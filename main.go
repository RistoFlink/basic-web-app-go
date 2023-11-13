package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

// in order for a function to response to a request from a browser, it needs to handle 2 parameters: ResponseWriter and Request
// in Go it is a custom to begin comments with the name of the function
// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")

}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}

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
