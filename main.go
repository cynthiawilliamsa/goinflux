package main

import (
	"fmt"
	"log"
	"net/http"
)

//Version in semantic versioning format
const version = "v0.0.1"

//Port Env variable for port connection
const Port = 8000

type server struct {
	router *http.ServeMux
}

func newServer() *server {
	s := &server{}
	s.router = http.DefaultServeMux
	s.routes()
	return s
}

func main() {
	server := newServer()
	log.Printf("service listening on port %d", Port)
	Address := fmt.Sprintf(":%d", Port)
	if err := http.ListenAndServe(Address, server.router); err != nil {
		log.Printf("server startup failed with error: %v\n, err", err)
	}
}
