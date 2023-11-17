package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RistoFlink/basic-web-app-go/pkg/config"
	"github.com/RistoFlink/basic-web-app-go/pkg/handlers"
	"github.com/RistoFlink/basic-web-app-go/pkg/render"
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
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//	http.HandleFunc("/", handlers.Repo.Home)
	//	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting the server at: %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
