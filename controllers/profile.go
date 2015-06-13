package controllers

import (
	"fmt"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// Display own profile if no id found
	if id == "" {
		// TODO
	}

	fmt.Fprintf(w, "You are looking at profile id %s", id)
}
