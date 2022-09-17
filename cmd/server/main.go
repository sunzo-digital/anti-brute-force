package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	handler := &Handler{}

	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server at :8080")
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server stopped")
}
