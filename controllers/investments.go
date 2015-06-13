package controllers

import (
	"fmt"
	"net/http"
)

func Investments(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are looking at the investments page")
}
