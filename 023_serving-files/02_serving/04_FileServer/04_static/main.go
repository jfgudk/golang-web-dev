package main

import (
	"log"
	"net/http"
)

func main() {
	// serves everything in this directory
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
