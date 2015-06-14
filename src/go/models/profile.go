package models

import (
	"appengine"
	"appengine/datastore"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type Profile struct {
	Name       string
	Bio        string `datastore:",noindex"`
	ProfilePic string
	Wallpaper  string
	Badges     []string
	Projects   []string
	Creditor   bool
}

// Each profile gets assigend the same ancestor so we have faster reads
func parentProfile(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Profile", "parent-profile", 0, nil)
}

func (p *Profile) generateKey(c appengine.Context, id string) (*datastore.Key,
	error) {
	key := datastore.NewKey(c, "Profile", id, 0, parentProfile(c))
	tmp_p := new(Profile)

	err := datastore.Get(c, key, tmp_p)
	if err == nil {
		return nil, errors.New("Name already in use")
	}

	return key, nil
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

func GetProfiles(r *http.Request) ([]Profile, error) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Profile").Ancestor(parentProfile(c)).Limit(3)

	var profiles []Profile
	_, err := q.GetAll(c, &profiles)
	return profiles, err
}

func (p *Profile) Store(r *http.Request) (string, error) {
	c := appengine.NewContext(r)
	id := url.QueryEscape(p.Name)

	key, err := p.generateKey(c, id)
	if err != nil {
		return "", err
	}

	_, err = datastore.Put(c, key, p)
	if err != nil {
		return "", err
	}

	return key.StringID(), nil
}
