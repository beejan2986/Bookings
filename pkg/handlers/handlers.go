package handlers

import (
	"net/http"

	"github.com/beejan2986/Bookings/pkg/config"
	"github.com/beejan2986/Bookings/pkg/models"
	"github.com/beejan2986/Bookings/pkg/render"
)

// Repository the repository use by handler
var Repo *Repository

// Repository is the Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "index.page.html", nil)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap})
}
