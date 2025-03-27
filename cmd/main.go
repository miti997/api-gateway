package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Gateway is running...")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("API Gateway is starting...")

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
