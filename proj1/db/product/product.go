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
	weight_gr   float32
	include_ts  string
	update_ts   string
}

func Find(dbCon *sql.DB, query string) ([]Product, error) {
	var products []Product
	rows, err := dbCon.Query(query)

	if err != nil {
		return nil, fmt.Errorf("error find product. query: %q, error: %v", query, err)
	}

	defer rows.Close()

	for rows.Next() {
		var p Product

		if err := rows.Scan(&p.id, &p.active, &p.name, &p.description, &p.xml_id, &p.width_mm, &p.height_mm, &p.weight_gr, &p.include_ts, &p.update_ts); err != nil {
			return nil, fmt.Errorf("Find product. query: %q; error: %v", query, err)
		}

		products = append(products, p)
	}

	return products, nil
}
