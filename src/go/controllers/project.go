package controllers

import (
	"net/http"
	"src/go/models"
	"strings"
)

func NewProject(w http.ResponseWriter, r *http.Request) {
	p := models.Project{
		Name:        "Hello World again",
		Description: "This is a test",
	}

	id, err := p.Store(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/project?id="+id, 307)
}

func Project(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")

	// Redirect wrong urls
	if len(url) != 3 {
		http.Redirect(w, r, "/404", 301)
		return
	}

	p, err := models.GetProject(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := getTemplate("project")
	varmap := map[string]interface{}{
		"id": p.Name,
	}

	render(t, w, varmap)
}
