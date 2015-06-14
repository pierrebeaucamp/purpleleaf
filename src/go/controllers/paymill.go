package controllers

import (
	"appengine"
	"appengine/urlfetch"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Paymill(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	amount := "100"
	currency := "EUR"
	token := "098f6bcd4621d373cade4e832627b4f6"
	description := "Your test transaction with a valid paymentobject ID"

	values := "amount=" + amount + "&currency=" + currency + "&token=" + token +
		"&description" + description

	data := bytes.NewReader([]byte(values))

	req, err := http.NewRequest("POST",
		"https://api.paymill.com/v2.1/transactions", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// private key
	req.SetBasicAuth("f322c3dc313539a55e3c0e0080e890dc", "")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := urlfetch.Client(c)
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(body))
}
