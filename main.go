package main

import (
	"net/http"

	"github.com/store-web-go/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
