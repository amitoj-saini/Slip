package database

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func Open_Database(cwd string) *sql.DB {
	db_filepath := filepath.Join(cwd, "db", "slip.db")
	if _, err := os.Stat(db_filepath); errors.Is(err, os.ErrNotExist) {
		os.Create(db_filepath)
	}

	db, err := sql.Open("sqlite3", db_filepath)
	if err != nil {
		log.Fatal("Could not open db!")
	}

	return db
}