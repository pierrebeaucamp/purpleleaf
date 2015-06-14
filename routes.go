package purpleleaf

import (
	"net/http"
	"src/go/controllers"
)

func init() {
	http.HandleFunc("/", controllers.Index)

	http.HandleFunc("/paymill", controllers.Paymill)

	http.HandleFunc("/project/", controllers.Project)
	http.HandleFunc("/project/new", controllers.SubmitProject)
	http.HandleFunc("/project/upload", controllers.NewProject)

	http.HandleFunc("/profile/", controllers.Profile)
	http.HandleFunc("/profile/dummies", controllers.Dummies)
}
