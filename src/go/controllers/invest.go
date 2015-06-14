package controllers

import (
	"appengine"
	"appengine/user"
	"net/http"
	"src/go/models"
)

func Invest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		http.Redirect(w, r, getLogin(r), 307)
	}

	p, err := models.GetProject(r)
	if err != nil {
		http.Redirect(w, r, "/", 307)
		return
	}

	t := getTemplate("invest")
	varmap := map[string]interface{}{
		"project": p,
		"login":   getLogin(r),
	}

	render(t, w, varmap)
}
