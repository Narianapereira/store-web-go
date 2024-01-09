package models

import (
	"github.com/store-web-go/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {

	database := db.ConnectDb()

	selectAllProducts, err := database.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Quantity = quantity
		p.Price = price

		products = append(products, p)
	}
	defer database.Close()
	return products
}