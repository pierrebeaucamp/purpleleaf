package models

import (
	"appengine"
	"appengine/blobstore"
	"appengine/datastore"
	"appengine/image"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type Project struct {
	Name        string
	Description string
	ImageURL    string
}

// Each project gets assigend the same ancestor so we have faster reads
func (p *Project) parent(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Project", "parent-project", 0, nil)
}

func GetImageURL(r *http.Request) string {
	c := appengine.NewContext(r)

	blobs, _, err := blobstore.ParseUpload(r)
	if err != nil {
		return "http://lorempixel.com/680/460/"
	}

	imgs := blobs["image"]
	if len(imgs) == 0 {
		return "http://lorempixel.com/680/460/"
	}

	o := image.ServingURLOptions{
		Secure: true,
	}

	img_url, err := image.ServingURL(c, imgs[0].BlobKey, &o)
	return img_url.String()
}

// generate a new key based on the title and return an error in case of
// duplicates
func (p *Project) generateKey(c appengine.Context, id string) (*datastore.Key,
	error) {
	key := datastore.NewKey(c, "Project", id, 0, p.parent(c))
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

	key := datastore.NewKey(c, "Project", id, 0, p.parent(c))

	err := datastore.Get(c, key, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func GetUploadURL(r *http.Request) (string, error) {
	c := appengine.NewContext(r)
	url, err := blobstore.UploadURL(c, "/project/upload", nil)
	return url.String(), err
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
