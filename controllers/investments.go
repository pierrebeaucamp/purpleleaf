package controllers

import (
	"net/http"
)

func Investments(w http.ResponseWriter, r *http.Request) {
	t := getTemplate("investments")
	render(t, w, nil)
}
