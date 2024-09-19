package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // _ is used to import a package only for its side effects
)

func Connect() (*sql.DB, error) {
	connectionString := "devbook_api:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Close(db *sql.DB) {
	db.Close()
}
