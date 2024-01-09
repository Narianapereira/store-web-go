package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	conection := "user=postgres dbname=my_store password=37141719nrp host=172.17.0.2 sslmode=disable"
	db, err := sql.Open("postgres", conection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
