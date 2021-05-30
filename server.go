package main

import (
	"log"
	"net/http"
	"strconv"
)

const (
	hdrContentType   = "Content-Type"
	hdrContentLength = "Content-Length"

	mimeJSON = "application/json"
)

//
func root(w http.ResponseWriter, req *http.Request) {
	data := []byte("Welcome.")
	contentLen := strconv.Itoa(len(data))
	w.Header().Set(hdrContentLength, contentLen)
	w.Header().Set(hdrContentType, mimeJSON)

	w.Write(data)
}

// startServer sets up handlers and starts server
func startServer(c *cfg) {
	log.Print("Starting server.")

	http.HandleFunc("/", root)

	// start
	log.Printf("Listing at '%s'.", c.address+":"+c.port)
	http.ListenAndServe(c.address+":"+c.port, nil)
}
