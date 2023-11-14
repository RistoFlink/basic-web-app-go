package main

import (
	"fmt"
	"net/http"

	"github.com/RistoFlink/basic-web-app-go/pkg/handlers"
)

const portNumber = ":8080"

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
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting the server at: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
