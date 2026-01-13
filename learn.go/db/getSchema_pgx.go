package main

// Util to check availability to connect PostgreSQL DB with jackc/pgx

import (
	"context"
	"fmt"

	pgx "github.com/jackc/pgx/v5"
)

func main() {
	url := "postgres://postgres:mysecretpassword@localhost:5432/postgres"

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Printf("Unable to connect DB <%s>. Error: %s", "postgres", err)
		return
	}
	defer conn.Close(context.Background())

	q := `SELECT table_name FROM information_schema.tables WHERE table_schema='public' ORDER BY table_name`
	rows, err := conn.Query(context.Background(), q)
	if err != nil {
		fmt.Printf("Unable to select from DB")
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Scan err:", err)
			return
		}
		fmt.Println("+T", name)
	}
	defer rows.Close()
}
