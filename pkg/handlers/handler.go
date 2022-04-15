package handlers

import (
	"github.com/MateerialRepo/go-webapp/pkg/config"
	"github.com/MateerialRepo/go-webapp/pkg/models"
	"github.com/MateerialRepo/go-webapp/pkg/render"
	"net/http"
)

//we would be using the repository pattern

//Repository is the repository data type
type Repository struct {
	App *config.AppConfig
}

//Repo is the repository used by the handler
var Repo *Repository

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home handles the homepage route of our application
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About handles the request to our about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	//send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
