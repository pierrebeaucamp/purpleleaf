package controllers

import (
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// Redirect to own profile if no id found and logged in
	if id == "" {
		// TODO
		http.Redirect(w, r, "/profile?id=123", 301)
	}

	t := getTemplate("profile")
	varmap := map[string]interface{}{
		"id": id,
	}

	render(t, w, varmap)
}
