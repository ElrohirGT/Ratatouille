package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DBDriver struct {
	Driver *sql.DB
}

func CreateDBDriver() {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}
}
