package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/store-web-go/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println(err.Error())
		}
		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			fmt.Println(err.Error())
		}

		models.CreateNewProduct(name, description, convertedPrice, convertedQuantity)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.Delete(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.GetProduct(productId)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do Id para int", err.Error())
		}
		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println(err.Error())
		}
		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			fmt.Println(err.Error())
		}
		models.Update(convertedId, convertedPrice, convertedQuantity, name, description)
	}
	http.Redirect(w, r, "/", 301)

}
