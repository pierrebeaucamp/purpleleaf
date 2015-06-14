package controllers

import (
	"appengine"
	"appengine/user"
	"html/template"
	"net/http"
	"net/url"
	"src/go/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// 404 page
	if r.URL.Path != "/" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}

	profiles, err := models.GetProfiles(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	projects, err := models.GetProjects(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t := getTemplate("index")
	varmap := map[string]interface{}{
		"profiles": profiles,
		"projects": projects,
		"login":    getLogin(r),
	}

	render(t, w, varmap)
}

func getLogin(r *http.Request) string {
	c := appengine.NewContext(r)
	u := user.Current(c)
	href := "/profile/"

	if u == nil {
		href, _ = user.LoginURL(c, r.URL.String())
	}

	return href
}

func getTemplate(name string) *template.Template {
	file := "src/go/views/" + name + ".html"

	funcmap := map[string]interface{}{
		"escape": url.QueryEscape,
	}

	return template.Must(template.New("").Funcs(funcmap).ParseFiles(
		file, "src/go/views/base.html"))
}

func render(t *template.Template, w http.ResponseWriter,
	varmap map[string]interface{}) {
	err := t.ExecuteTemplate(w, "body", varmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
