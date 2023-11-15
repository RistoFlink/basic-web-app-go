package handlers

import (
	"net/http"

	"github.com/RistoFlink/basic-web-app-go/pkg/config"
	"github.com/RistoFlink/basic-web-app-go/pkg/render"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo is the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// in order for a function to response to a request from a browser, it needs to handle 2 parameters: ResponseWriter and Request
// in Go it is a custom to begin comments with the name of the function
// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")

}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
