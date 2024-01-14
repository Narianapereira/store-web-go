package models

import (
	"log"

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

	selectAllProducts, err := database.Query("select * from products order by id asc")
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

func GetProduct(idToGet string) Product {
	db := db.ConnectDb()
	productData, err := db.Query("select * from products where id=$1", idToGet)
	if err != nil {
		panic(err.Error())
	}

	productReturn := Product{}
	var id, quantity int
	var name, description string
	var price float64

	for productData.Next() {
		err = productData.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		productReturn.Id = id
		productReturn.Name = name
		productReturn.Description = description
		productReturn.Price = price
		productReturn.Quantity = quantity
	}
	defer db.Close()
	return productReturn
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

func Update(id int, price float64, quantity int, name string, description string) {
	db := db.ConnectDb()

	updateData, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		log.Println(err.Error())
	}

	updateData.Exec(name, description, price, quantity, id)
	defer db.Close()
}
