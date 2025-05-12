package main

import (
	"database/sql"
	"os"
	"proj1/conf"
	"proj1/db"
	"proj1/db/product"
)

var dbCon *sql.DB

func main() {
	conf.SetDbAccess()

	dbCon := db.Connect(
		os.Getenv("USER"),
		os.Getenv("PASSWD"),
		os.Getenv("NET_TYPE"),
		os.Getenv("HOST_PORT"),
		os.Getenv("DB_NAME"),
	)

	products := product.Find(dbCon, "SELECT * FROM product")

	/*
		- получить тестовые товары
		- вывести их в api
		- выввести на фронт в js
	*/
}
