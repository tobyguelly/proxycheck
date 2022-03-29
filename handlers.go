package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	w.WriteHeader(200)
	_, _ = w.Write([]byte("Hello World!"))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	query := r.URL.Query()
	codes, present := query["code"]
	if !present || len(codes) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Please provide the query parameter code with a desired status code!"))
	} else {
		code, err := strconv.Atoi(codes[0])
		if err != nil {
			code = http.StatusBadRequest
		}
		w.WriteHeader(code)
		_, _ = w.Write([]byte(http.StatusText(code)))
	}
}

func headersHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	_, _ = fmt.Fprintf(w, "%s %s %s\n", r.Method, r.RequestURI, r.Proto)
	_, _ = fmt.Fprintf(w, "%s: %s\n", "Host", r.Host)
	for key, value := range r.Header {
		_, _ = fmt.Fprintf(w, "%s: %s\n", key, strings.Join(value, "; "))
	}
}

func ignoreHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
}
