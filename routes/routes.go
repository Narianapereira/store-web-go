package routes

import (
	"net/http"

	"github.com/store-web-go/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)

}
