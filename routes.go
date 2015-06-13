package purpleleaf

import (
	"net/http"

	"controllers"
)

func init() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/project", controllers.Project)
	http.HandleFunc("/profile", controllers.Profile)
	http.HandleFunc("/investments", controllers.Investments)
}
