package controllers

import (
	"html/template"
	"net/http"
)

func Investments(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/investments.html",
		"views/base.html"))

	err := t.ExecuteTemplate(w, "body", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
