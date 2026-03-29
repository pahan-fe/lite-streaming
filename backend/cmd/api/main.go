package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	fmt.Println("Starting API server on port 8080...")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))  
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}