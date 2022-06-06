package database

import (
	"database/sql"
	"os"
)

func createConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	return db, nil
}
