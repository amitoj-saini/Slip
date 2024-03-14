package main

import (
	"log"
	"net/http"
	"slip/config"
	"slip/internal/handlers"
	"slip/internal/middleware"
)

func main() {
	config.LoadEnv()

	// server configs
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// routes
	http.HandleFunc("/", middleware.AuthHandler(handlers.Home))

	// find addr user wants to host serv on
	addr := config.GetEnv("ADDR", ":8080")

	log.Println("Http server started on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
