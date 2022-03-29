package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/headers", headersHandler)
	http.HandleFunc("/ignore", ignoreHandler)
	log.Printf("proxycheck server launching...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Failed to start server: %s\n", err.Error())
	}
}
