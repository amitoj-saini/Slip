package database

import (
	"database/sql"
	"log"
)

func Initialize_Models(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			email UNIQUE
		)
	`)

	if err != nil {
		log.Fatal("Could not create models!")
	}
}