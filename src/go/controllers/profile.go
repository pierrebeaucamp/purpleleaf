package controllers

import (
	"net/http"
	"src/go/models"
	"strings"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")

	// Redirect to own profile if no id found and logged in
	if len(url) != 3 {
		// TODO
		http.Redirect(w, r, "/profile/otto", 307)
		return
	}

	p, err := models.GetProfile(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := getTemplate("profile")
	varmap := map[string]interface{}{
		"id": p.Name,
	}

	render(t, w, varmap)
}
