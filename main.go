package main

import (
	"os"
	"log"
	"net/http"
	"slip/config"
	"slip/internal/handlers"
	"slip/internal/middleware"
	"slip/internal/database"
)

func main() {
	// configurations
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not get CWD!")
	}
	config.LoadEnv()

	db := database.Open_Database(cwd)
	database.Initialize_Models(db)

	// server configs
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// routes
	http.HandleFunc("/", middleware.AuthHandler(handlers.Home, db))

	// find addr user wants to host serv on
	addr := config.GetEnv("ADDR", ":8080")

	log.Println("Http server started on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
