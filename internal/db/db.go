package db

import (
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Connect() error {
	var err error
	db, err = sql.Open("sqlite3", "localft.sqlite")
	if err != nil {return err}

	schema := `
	CREATE TABLE IF NOT EXISTS files (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		extension TEXT NOT NULL,
		mime_type TEXT NOT NULL,
		size INTEGER,
		created_at INTEGER NOT NULL
	);`

	if _, err := db.Exec(schema); err != nil {
		return err
	}

	return db.Ping()
}

func GetDb() (*sql.DB, error) {
	if db == nil {
		err := Connect()
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
