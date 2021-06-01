package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	hdrContentType   = "Content-Type"
	hdrContentLength = "Content-Length"

	mimeJSON = "application/json"
)

//
func root(w http.ResponseWriter, req *http.Request) {
	data := []byte("Welcome.\n")
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
	s := &http.Server{
		Addr:           ":" + c.port,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Listing on port '%s'.", s.Addr)
	log.Fatal(s.ListenAndServe())
}
