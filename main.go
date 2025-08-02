package main

import (
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	ID          int
	Title       string
	Author      string
	Description string
	Price       float64
}



func main() {
	// Define HTTP handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	
	// Start HTTP server
	fmt.Println("Server starting on port 80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}
