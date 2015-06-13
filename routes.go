package purpleleaf

import (
	"net/http"

	"controllers"
)

func init() {
	http.HandleFunc("/", controllers.Index)
}
