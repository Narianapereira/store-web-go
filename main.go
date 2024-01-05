package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)

}

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camiseta", "Azul, bem bonita", 39, 5},
		{"Tenis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Produto novo", "Muito legal", 1.99, 1},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
