package middleware

import (
	"slip/internal/handlers"
	"database/sql"
	"slip/config"
	"net/http"
	"log"
)

func AuthHandler(next config.HandlerWithDB, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM USERS")
		if err != nil {
			log.Fatal("Could not execute query in db!")
		}
		defer rows.Close()

		users := 0
		for rows.Next() {
			users++
		}

		if users == 0 {
			handlers.GetStarted(w, r, db)
		} else {
			next(w, r, db)
		}
	}
}
