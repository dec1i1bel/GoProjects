package main

import (
	"os"
	"proj1/conf"
	"proj1/db"
)

type Product struct {
	active string
	name   string
}

func main() {
	conf.SetDbAccess()
	db.Connect(
		os.Getenv("USER"),
		os.Getenv("PASSWD"),
		os.Getenv("NET_TYPE"),
		os.Getenv("HOST_PORT"),
		os.Getenv("DB_NAME"),
	)

	/*
		- получить тестовые товары
		- вывести их в api
		- выввести на фронт в js
	*/
}
