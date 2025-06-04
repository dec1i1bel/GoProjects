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
		> POST добавление товара
			https://go.dev/doc/tutorial/web-service-gin#add_item
		- GET удаление товара
	- добавление пользователя по email и паролю
	- авторизация
		- отправка кода подтверждения на email
		- проверка кода
	- ограничение доступности роутов для неавторизованного
		- неавторизованному - получение максимум трёх товаров
		- авторизованному - все роуты без ограничений
	- у авторизованного - временная сессия по токену. по истечении срока - запрос повторной авторизации по api
	*/

	router := gin.Default()
	router.GET("/products", findProducts)
	router.GET("/product/:id", findProductById)
	// router.POST("/product/add", addProduct)

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

func addProduct(c *gin.Context) (int64, error) {
	var product Product
}
