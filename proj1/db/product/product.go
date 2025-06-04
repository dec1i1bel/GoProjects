package product

import (
	"database/sql"
	"fmt"
	"os"
	"proj1/db"
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

func CreateDbConnection() *sql.DB {
	return db.Connect(
		os.Getenv("USER"),
		os.Getenv("PASSWD"),
		os.Getenv("NET_TYPE"),
		os.Getenv("HOST_PORT"),
		os.Getenv("DB_NAME"),
	)
}
