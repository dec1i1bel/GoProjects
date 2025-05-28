package product

import (
	"database/sql"
	"fmt"
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

func FindAll(dbCon *sql.DB) ([]Product, error) {
	var products []Product

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
