package product

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"proj1/db"
	"time"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          int64
	Active      string
	Name        string
	Description string
	Xml_id      string
	Width_mm    float32
	Height_mm   float32
	Weight_gr   float32
}

// func FindAll(dbCon *sql.DB) ([]Product, error) {
func FindAll() ([]Product, error) {
	var products []Product
	dbCon := CreateDbConnection()
	rows, err := dbCon.Query("SELECT id, active, name, description, xml_id, width_mm, height_mm, weight_gr FROM product")

	if err != nil {
		fmt.Printf("Error find all products: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var p Product

		if err := rows.Scan(&p.Id, &p.Active, &p.Name, &p.Description, &p.Xml_id, &p.Width_mm, &p.Height_mm, &p.Weight_gr); err != nil {
			fmt.Printf("Error next product: %v", err)
		}

		products = append(products, p)
	}

	return products, nil
}

func FindById(id int64) (Product, error) {
	var product Product
	dbCon := CreateDbConnection()
	row := dbCon.QueryRow("SELECT id, active, name, description, xml_id, width_mm, height_mm, weight_gr FROM product WHERE id = ?", id)

	if err := row.Scan(&product.Id, &product.Active, &product.Name, &product.Description, &product.Xml_id, &product.Width_mm, &product.Height_mm, &product.Weight_gr); err != nil {
		if err == sql.ErrNoRows {
			return product, fmt.Errorf("product::findById %d: no such product", id)
		}

		return product, fmt.Errorf("product::findById %d: %v", id, err)
	}

	return product, nil
}

func Insert(c *gin.Context) (string, error) {
	var p Product

	if err := c.BindJSON(&p); err != nil {
		return "error", fmt.Errorf("Error binding new product json to Product object. Product: %q, error: %v", c.product, err)
	}

	dbCon := CreateDbConnection()

	// toDo: doc insert row in Go: https://golangbot.com/mysql-create-table-insert-row/#insert-row
	query := "INSERT INTO product (name, active, description, xml_id, width_mm, height_mm, weight_gr) VALUES (?, ?, ?, ?, ?, ?, ?)"
	// create a context with a timeout so that the query times out in case of any network, partition or runtime errors
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	// закроет соединение с БД
	defer cancelfunc()

	// create a prepared statement for the insert query using this template
	stmt, err := dbCon.PrepareContext(ctx, query)

	if err != nil {
		log.Printf("Error <%s> preparing SQL statement", err)
		return "error", err
	}

	// close statement after use
	defer stmt.Close()

	// The ExecContext method of the DB package executes any query that doesn’t return any rows
	res, err := stmt.ExecContext(ctx, p.Name, p.Active, p.Description, p.Xml_id, p.Height_mm, p.Weight_gr)

	if err != nil {
		log.Printf("Error executing context on inserting product: %s", err)
		return "error", err
	}

	row, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error getting rows affected on insering product: %s", err)
		return "error", err
	}

	log.Printf("%d product created", row)

	return "success", nil
}

func CreateDbConnection() *sql.DB {
	return db.Connect(
		os.Getenv("USER"),
		os.Getenv("PASSWD"),
		os.Getenv("NET_TYPE"),
		os.Getenv("HOST_PORT"),
		os.Getenv("DB_NAME"),
	)
}
