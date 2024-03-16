package handlers

import (
	"database/sql"
	"html/template"
	"slip/config"
	"net/http"
	"log"
)



func GetStarted(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	data := config.TemplateVaribles {
		Title: "Get Started",
	}
	
	if r.Method == http.MethodGet {
		tmpl, err := template.New("").ParseFiles("templates/get-started.html", "templates/template.html")
		if err != nil {
			log.Fatal("Could not find template!", err)
		}

		tmpl.ExecuteTemplate(w, "base", data)
	}
}
