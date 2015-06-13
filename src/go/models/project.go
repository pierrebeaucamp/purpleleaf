package models

import (
	"appengine"
	"appengine/datastore"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type Project struct {
	Name        string
	Description string
}

// Each project gets assigend the same ancestor so we have faster reads
func parentProject(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Project", "parent-project", 0, nil)
}

// generate a new key based on the title and return an error in case of
// duplicates
func (p *Project) generateKey(c appengine.Context, id string) (*datastore.Key,
	error) {
	key := datastore.NewKey(c, "Project", id, 0, parentProject(c))
	tmp_p := new(Project)

	err := datastore.Get(c, key, tmp_p)
	if err == nil {
		return nil, errors.New("Title already in use")
	}

	return key, nil
}

// returns a project based on an id
func GetProject(r *http.Request) (*Project, error) {
	c := appengine.NewContext(r)
	id := strings.Split(r.URL.Path, "/")[2]
	p := new(Project)

	key := datastore.NewKey(c, "Project", id, 0, parentProject(c))

	err := datastore.Get(c, key, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// stores a project
func (p *Project) Store(r *http.Request) (string, error) {
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
