package main

import (
	"fmt"
	"net/http"
	"os"
	"proj1/conf"
	"proj1/db"
	"proj1/db/product"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.SetDbAccess()

	/* toDo:
	+ получить тестовые товары
	+ вывести их в api
		+ перенести получение в метод, подключить его к роутеру
		+ в db\product корректно сформировать объект для json
		+ в app сконвертировать struct в json
	- получить на фронт в js
	*/

	router := gin.Default()
	router.GET("/products", findProducts)

	router.Run("localhost:8080")
}

func findProducts(context *gin.Context) {
	dbCon := db.Connect(
		os.Getenv("USER"),
		os.Getenv("PASSWD"),
		os.Getenv("NET_TYPE"),
		os.Getenv("HOST_PORT"),
		os.Getenv("DB_NAME"),
	)

	products, err := product.FindAll(dbCon)

	if err != nil {
		fmt.Printf("error getting products: %v", err)
	}

	context.JSON(http.StatusOK, products)
}
