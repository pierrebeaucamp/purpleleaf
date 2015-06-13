package controllers

import (
	"net/http"
	"src/go/models"
	"strings"
)

func NewProject(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("title")

	if name == "" {
		http.Redirect(w, r, "/", 307)
		return
	}

	p := models.Project{
		Name:        name,
		Description: r.FormValue("description"),
		ImageURL:    models.GetImageURL(r),
	}

	id, err := p.Store(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/project/"+id, 307)
}

func Project(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")

	// Redirect wrong urls
	if len(url) != 3 || url[2] == "" {
		http.Redirect(w, r, "/", 307)
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

func SubmitProject(w http.ResponseWriter, r *http.Request) {
	url, err := models.GetUploadURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t := getTemplate("projectForm")
	varmap := map[string]interface{}{
		"url": url,
	}

	render(t, w, varmap)
}
