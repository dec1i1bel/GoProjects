package main

import (
	"database/sql"
	"fmt"
	"log"
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

	/* toDo:
	+ получить тестовые товары
		+ тест
	- вывести их в api
	- выввести на фронт в js
	*/

	products, err := product.Find(dbCon, "SELECT * FROM product")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("success: all products found: %v\n", products)
}
