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

		p.Id = id
		p.Name = name
		p.Description = description
		p.Quantity = quantity
		p.Price = price

		products = append(products, p)
	}
	defer database.Close()
	return products
}

func CreateNewProduct(name, description string, convertedPrice float64, convertedQuantity int) {
	db := db.ConnectDb()
	insertData, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, convertedPrice, convertedQuantity)
	defer db.Close()
}

func Delete(productId string) {
	db := db.ConnectDb()
	deleteData, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deleteData.Exec(productId)
	defer db.Close()
}
