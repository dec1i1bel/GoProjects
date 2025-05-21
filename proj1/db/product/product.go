package product

import (
	"database/sql"
	"fmt"
)

type Product struct {
	id          int64
	active      string
	name        string
	description string
	xml_id      string
	width_mm    float32
	height_mm   float32
}

func Find(dbCon *sql.DB, query string) ([]Product, error) {
	var products []Product
	rows, err := dbCon.Query(query)

	if err != nil {
		return nil, fmt.Errorf("error find product. query: %q, error: %v", query, err)
	}

	defer rows.Close()

	for rows.Next() {
		var product Product

		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("Find product. query: %q; error: %v", query, err)
		}

		products = append(products, product)
	}

	return products, nil
}
