package controllers

import (
	"appengine"
	"appengine/user"
	"net/http"
	"net/url"
	"src/go/models"
	"strings"
)

func Dummies(w http.ResponseWriter, r *http.Request) {
	var err error

	p1 := models.Profile{
		Name:       "Herbert Hubel",
		ProfilePic: "http://lorempixel.com/680/460/",
		Wallpaper:  "http://lorempixel.com/680/460/",
		Bio:        "Lorem ipsum...",
		Badges:     []string{"health20", "ag5", "edu100", "trade50"},
		Projects:   []string{"Hello World", "Bye World"},
		Creditor:   true,
	}

	p2 := models.Profile{
		Name:       "Umar Halitnov",
		ProfilePic: "http://lorempixel.com/680/460/",
		Wallpaper:  "http://lorempixel.com/680/460/",
		Bio:        "Lorem ipsum...",
		Badges:     []string{"health20", "ag5", "edu100", "trade50"},
		Projects:   []string{"Hello World"},
		Creditor:   false,
	}

	p3 := models.Profile{
		Name:       "Family Doe",
		ProfilePic: "http://lorempixel.com/680/460/",
		Wallpaper:  "http://lorempixel.com/680/460/",
		Bio:        "Lorem ipsum...",
		Badges:     []string{"health20", "ag5", "edu100", "trade50"},
		Projects:   []string{"Bye World"},
		Creditor:   false,
	}

	for _, p := range [...]models.Profile{p1, p2, p3} {
		_, err = p.Store(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	pr1 := models.Project{
		Name:         "Hello World",
		Description:  "lorem ipsum...",
		ImageURL:     "http://lorempixel.com/680/460/",
		WallpaperURL: "http://lorempixel.com/680/460/",
		Investors:    []string{"Herbert Hubel"},
		Goal:         1000,
		Raised:       1000,
	}

	pr2 := models.Project{
		Name:         "Bye World",
		Description:  "lorem ipsum...",
		ImageURL:     "http://lorempixel.com/680/460/",
		WallpaperURL: "http://lorempixel.com/680/460/",
		Investors:    []string{"Herbert Hubel"},
		Goal:         1000,
		Raised:       500,
	}

	for _, p := range [...]models.Project{pr1, pr2} {
		_, err = p.Store(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	http.Redirect(w, r, "/", 307)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	// Redirect to own profile if no id found and logged in
	if len(path) != 3 || path[2] == "" {
		c := appengine.NewContext(r)
		u := user.Current(c)

		if u == nil {
			http.Redirect(w, r, "/", 307)
		}

		id := url.QueryEscape(u.String())
		http.Redirect(w, r, "/profile/"+id, 302)
	}

	p, err := models.GetProfile(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var badges = make([]models.Badge, len(p.Badges))
	for i, b := range p.Badges {
		badges[i] = models.GetBadge(b)
	}

	var projects = make([]models.Project, len(p.Projects))
	for i, pr := range p.Projects {
		tmp, _ := models.GetProjectById(r, url.QueryEscape(pr))
		projects[i] = *tmp
	}

	t := getTemplate("profile")
	varmap := map[string]interface{}{
		"profile":  p,
		"badges":   badges,
		"projects": projects,
		"login":    getLogin(r),
	}

	render(t, w, varmap)
}
