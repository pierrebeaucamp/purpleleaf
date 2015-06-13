package purpleleaf

import (
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/base.html"))

	err := t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
