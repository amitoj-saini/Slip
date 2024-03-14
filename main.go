package main

import (
	"net/http"
	"log"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	addr := os.Getenv("ADDR")

	log.Println("Http server started on ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
