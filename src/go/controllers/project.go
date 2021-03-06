package controllers

import (
	"appengine"
	"appengine/user"
	"net/http"
	"net/url"
	"src/go/models"
	"strings"
)

func NewProject(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("title")
	if name == "" {
		http.Redirect(w, r, "/", 307)
	}

	p := models.Project{
		Name:         name,
		Description:  r.FormValue("description"),
		ImageURL:     models.GetImageURL(r),
		WallpaperURL: "http://lorempixel.com/680/460/",
	}

	id, err := p.Store(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/project/"+id, 307)
}

func Project(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	// Redirect wrong urls
	if len(path) != 3 || path[2] == "" {
		http.Redirect(w, r, "/", 307)
	}

	if r.URL.Query().Get("action") == "invest" {
		Invest(w, r)
		return
	}

	p, err := models.GetProject(r)
	if err != nil {
		http.Redirect(w, r, "/", 307)
	}

	var profiles = make([]models.Profile, len(p.Investors))
	for i, inv := range p.Investors {
		tmp, _ := models.GetProfileById(r, url.QueryEscape(inv))
		profiles[i] = *tmp
	}

	t := getTemplate("project")
	varmap := map[string]interface{}{
		"project":  p,
		"profiles": profiles,
		"login":    getLogin(r),
	}

	render(t, w, varmap)
}

func SubmitProject(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		http.Redirect(w, r, getLogin(r), 307)
	}

	url, err := models.GetUploadURL(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := getTemplate("projectForm")
	varmap := map[string]interface{}{
		"url":   url,
		"login": getLogin(r),
	}

	render(t, w, varmap)
}
