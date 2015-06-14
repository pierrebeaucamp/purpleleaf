package controllers

import (
	"appengine"
	"appengine/user"
	"net/http"
	"net/url"
	"src/go/models"
	"strings"
)

func Dummies(w http.ResponseWriter, r *http.Request) {
	p1 := models.Profile{
		Name:       "Herbert Hubel",
		ProfilePic: "http://lorempixel.com/680/460/",
	}

	p2 := models.Profile{
		Name:       "Umar Halitnov",
		ProfilePic: "http://lorempixel.com/680/460/",
	}

	p3 := models.Profile{
		Name:       "Family Doe",
		ProfilePic: "http://lorempixel.com/680/460/",
	}

	var id string
	var err error

	for _, p := range [...]models.Profile{p1, p2, p3} {
		id, err = p.Store(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	http.Redirect(w, r, "/profile/"+id, 307)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	// Redirect to own profile if no id found and logged in
	if len(path) != 3 || path[2] == "" {
		c := appengine.NewContext(r)
		u := user.Current(c)

		if u == nil {
			http.Redirect(w, r, "/", 307)
		}

		id := url.QueryEscape(u.String())
		http.Redirect(w, r, "/profile/"+id, 302)
	}

	p, err := models.GetProfile(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := getTemplate("profile")
	varmap := map[string]interface{}{
		"id":    p.Name,
		"login": getLogin(r),
	}

	render(t, w, varmap)
}
