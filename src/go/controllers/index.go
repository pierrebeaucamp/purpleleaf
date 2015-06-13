package controllers

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// 404 page
	if r.URL.Path != "/" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}

	t := getTemplate("index")
	render(t, w, nil)
}

func getTemplate(name string) *template.Template {
	file := "src/go/views/" + name + ".html"
	return template.Must(template.ParseFiles(file, "src/go/views/base.html"))
}

func render(t *template.Template, w http.ResponseWriter,
	varmap map[string]interface{}) {
	err := t.ExecuteTemplate(w, "body", varmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
