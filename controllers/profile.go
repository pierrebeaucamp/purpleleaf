package controllers

import (
	"html/template"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// Redirect to own profile if no id found and logged in
	if id == "" {
		// TODO
		http.Redirect(w, r, "/profile?id=123", 301)
	}

	t := template.Must(template.ParseFiles("views/profile.html",
		"views/base.html"))

	varmap := map[string]interface{}{
		"id": id,
	}

	err := t.ExecuteTemplate(w, "body", varmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
