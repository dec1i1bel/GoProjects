package product

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"proj1/db"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          *int64   `json:"id,omitempty"`
	Active      string   `json:"active"`
	Name        string   `json:"name"`
	Description *string  `json:"description,omitempty"`
	Xml_id      *string  `json:"xml_id,omitempty"`
	Width_mm    *float32 `json:"width_mm,omitempty"`
	Height_mm   *float32 `json:"height_mm,omitempty"`
	Weight_gr   *float32 `json:"weight_mm,omitempty"`
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
		return "error binding json to product", err
	}

	sqlFields := "active, name"
	var sqlValues string
	var params []interface{}
	params = append(params, p.Active)
	params = append(params, p.Name)

	if p.Description != nil {
		sqlFields += ", description"
		params = append(params, p.Description)
	}

	if p.Xml_id != nil {
		sqlFields += ", xml_id"
		params = append(params, p.Xml_id)
	}

	if p.Width_mm != nil {
		sqlFields += ", width_mm"
		params = append(params, p.Width_mm)
	}

	if p.Height_mm != nil {
		sqlFields += ", height_mm"
		params = append(params, p.Height_mm)
	}

	if p.Weight_gr != nil {
		sqlFields += ", weight_gr"
		params = append(params, p.Weight_gr)
	}

	fieldsSpl := strings.Split(sqlFields, ", ")

	for i := 0; i < len(fieldsSpl); i++ {
		sqlValues += "?, "
	}

	sqlValues = strings.Trim(sqlValues, ", ")

	fmt.Println(sqlFields)
	fmt.Println(sqlValues)

	dbCon := CreateDbConnection()

	query := "INSERT INTO product (" + sqlFields + ") VALUES (" + sqlValues + ")"
	// fmt.Println(query)
	// create a context with a timeout so that the query times out in case of any network, partition or runtime errors
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	// close db connection
	defer cancelfunc()

	// create a prepared statement for the query using the template
	stmt, err := dbCon.PrepareContext(ctx, query)

	if err != nil {
		return "error preparing sql statement", err
	}

	// close statement after use
	defer stmt.Close()

	// The ExecContext method of the DB package executes any query that doesn’t return any rows
	res, err := stmt.ExecContext(ctx, params...)

	if err != nil {
		log.Printf("Error inserting product: %s", err)
		return "error inserting product", err
	}

	row, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error getting rows affected on insering product: %s", err)
		return "error getting rows affected on insering product", err
	}

	log.Printf("%d ==product created==", row)

	return "success", nil
}

func Delete(id int64) (string, error) {
	dbCon := CreateDbConnection()
	query := "DELETE FROM product WHERE id=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := dbCon.PrepareContext(ctx, query)

	if err != nil {
		return "error preparing sql statement", err
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, id)

	if err != nil {
		log.Printf("Error deleting product: %s", err)
		return "error", err
	}

	row, err := res.RowsAffected()

	if err != nil {
		log.Printf("Error getting rows affected on deleting product: %s", err)
		return "error", err
	}

	log.Printf("%d ==product deleted==", row)

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
