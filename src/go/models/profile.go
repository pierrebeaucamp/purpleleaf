package models

import (
	"appengine"
	"appengine/datastore"
	"net/http"
	"strings"
)

type Profile struct {
	Name       string
	ProfilePic string
}

// Each profile gets assigend the same ancestor so we have faster reads
func parentProfile(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Profile", "parent-profile", 0, nil)
}

// returns a profile based on an id
func GetProfile(r *http.Request) (*Profile, error) {
	c := appengine.NewContext(r)
	id := strings.Split(r.URL.Path, "/")[2]
	p := new(Profile)

	key := datastore.NewKey(c, "Profile", id, 0, parentProfile(c))

	err := datastore.Get(c, key, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
