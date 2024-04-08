package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDriver(user, password, dbname, host, port string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user,
		password,
		dbname,
		host,
		port,
	)
	db, err := sql.Open("postgres", connStr)
	return db, err
}
