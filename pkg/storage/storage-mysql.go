package storage

import "database/sql"

type Storage struct {
	db *sql.DB
}

func SetUpStorage() (*Storage, error) {
	db, err = sql.Open("mysql", "")
}
