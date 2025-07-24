package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB(path string) *sql.DB {
    db, err := sql.Open("sqlite3", path)
    if err != nil {
        log.Fatal(err)
    }
    return db
}