package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDriver(user, password, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=verify-full", user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	return db, err
}
