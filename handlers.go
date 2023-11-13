package main

import (
	"net/http"
)

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