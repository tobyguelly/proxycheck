package main

import (
	"log"
	"net/http"
)

func logRequest(r *http.Request) {
	log.Printf("[%s] %s %s %s\n", r.RemoteAddr, r.Method, r.RequestURI, r.Proto)
}
