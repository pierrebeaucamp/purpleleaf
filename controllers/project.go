package controllers

import (
	"html/template"
	"net/http"
)

func Project(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// Redirect to index if no id found
	if id == "" {
		http.Redirect(w, r, "/", 302)
	}

	t := template.Must(template.ParseFiles("views/project.html",
		"views/base.html"))

	varmap := map[string]interface{}{
		"id": id,
	}

	err := t.ExecuteTemplate(w, "body", varmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
