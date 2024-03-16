package handlers

import (
	"database/sql"
	"net/http"
	"fmt"
)

func Home(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Fprintf(w, "Welcome to the home page!")
}
