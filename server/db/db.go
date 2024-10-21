package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func openConnection() (*sql.DB, error) {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	return sql.Open("postgres", connStr)
}