package db

import (
	"database/sql"
	_ "github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

var defaultPgUrl = "postgres://postgres@127.0.0.1:5432/badge?sslmode=disable"

func Connect(connStr *string) (*sql.DB, error) {
	if connStr == nil {
		connStr = &defaultPgUrl
	}

	db, err := sql.Open("postgres", *connStr)
	return db, err
}

func CreateTables(conn *sql.DB) error {
	err := CreateCategoryTable(conn)
	if err != nil {
		return err
	}

	err = CreateUserTable(conn)
	if err != nil {
		return err
	}

	err = CreateMessageTable(conn)
	if err != nil {
		return err
	}

	return nil
}

