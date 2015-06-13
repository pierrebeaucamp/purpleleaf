package controllers

import (
	"net/http"
)

func Project(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// Redirect to index if no id found
	if id == "" {
		http.Redirect(w, r, "/", 302)
	}

	t := getTemplate("project")
	varmap := map[string]interface{}{
		"id": id,
	}

	render(t, w, varmap)
}
