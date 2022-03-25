package handlers

import (
	"github.com/williamneo/TS/webapp/pkg/config"
	"github.com/williamneo/TS/webapp/pkg/models"
	"github.com/williamneo/TS/webapp/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	app *config.Appconfig
}

//NewRepo create a new repository
func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		app: a,
	}
}

//NewHandler set the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again!"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
