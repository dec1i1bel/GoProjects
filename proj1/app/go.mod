module proj1/app

go 1.24.2

replace proj1/db => ../db
replace proj1/conf => ../conf

require proj1/db v0.0.0-00010101000000-000000000000
require proj1/conf v0.0.0-00010101000000-000000000000

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.2 // indirect
)
