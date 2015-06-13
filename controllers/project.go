package controllers

import (
	"fmt"
	"net/http"
)

func Project(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// Redirect to index if no id found
	if id == "" {
		http.Redirect(w, r, "/", 302)
	}

	fmt.Fprintf(w, "You are looking at project id %s", id)
}
