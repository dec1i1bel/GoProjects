module proj1/db/product

go 1.24.2

replace proj1/db => ../

require proj1/db v0.0.0-00010101000000-000000000000

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.2 // indirect
)
