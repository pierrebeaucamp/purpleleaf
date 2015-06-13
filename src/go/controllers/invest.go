package controllers

import (
	"net/http"
)

func Invest(w http.ResponseWriter, r *http.Request) {
	t := getTemplate("invest")
	render(t, w, nil)
}
