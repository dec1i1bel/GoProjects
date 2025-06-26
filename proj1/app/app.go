package main

import (
	"fmt"
	"net/http"
	"proj1/conf"
	"proj1/db/product"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.SetDbAccess()

	/* toDo - api:
	> роуты:
		+ все товары
		+ товар по id
		+ POST добавление товара с минимальным набором полей (активность, название)
			+ код
			+ тест
		+ GET удаление товара
			+ код
			+ тест
		- POST добавление товара с произвольным набором полей в запросе
			- добавление товара
				- тест на корректных данных в запросе
			- валидация полей
				- тест на некорректнеых
	- добавление пользователя по email и паролю
	- авторизация
		- отправка кода подтверждения на email
		- проверка кода
	- ограничение доступности роутов для неавторизованного
		- неавторизованному - получение максимум трёх товаров
		- авторизованному - все роуты без ограничений
	- у авторизованного - временная сессия по токену. по истечении срока - запрос повторной авторизации по api

	примеры запросов curl:

	curl http://localhost:8080/product/add -X POST -H 'Content-Type: application/json' -d '{"name":"name1"}'
	curl http://localhost:8080/product/add -X POST -H 'Content-Type: application/json' -d '{"active":"y", "name":"name1"}'
	curl http://localhost:8080/product/delete/1

	*/

	router := gin.Default()
	router.GET("/products", findProducts)
	router.GET("/product/:id", findProductById)
	router.GET("/product/delete/:id", deleteProduct)
	router.POST("/product/add", insertProduct)

	router.Run("localhost:8080")
}

func findProducts(c *gin.Context) {
	products, err := product.FindAll()

	if err != nil {
		fmt.Printf("error getting products: %v", err)
	}

	c.JSON(http.StatusOK, products)
}

func findProductById(c *gin.Context) {
	idRaw := strings.TrimSpace(c.Param("id"))
	id, err := strconv.ParseInt(idRaw, 10, 64)

	if err != nil {
		fmt.Printf("error converting param id <%v> to int\n", idRaw)
	}

	product, err := product.FindById(id)

	if err != nil {
		fmt.Printf("error getting product by id %d\n", id)
	}

	c.JSON(http.StatusOK, product)
}

func insertProduct(c *gin.Context) {
	newProductId, err := product.Insert(c)

	if err != nil {
		fmt.Printf("error add product: %v", c.Param("product"))
	}

	c.JSON(http.StatusCreated, newProductId)
}

func deleteProduct(c *gin.Context) {
	idRaw := strings.TrimSpace(c.Param("id"))
	id, err := strconv.ParseInt(idRaw, 10, 64)

	if err != nil {
		fmt.Printf("error converting param id <%v> to int\n", idRaw)
	}

	deleteProductId, err := product.Delete(id)

	if err != nil {
		fmt.Printf("Error on delete product: %v", c.Param("product"))
	}

	c.JSON(http.StatusAccepted, deleteProductId)
}
