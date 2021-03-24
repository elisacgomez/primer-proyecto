package main

import (
	"log"
	"net/http"
	"primer-proyecto/pkg/server"
)

func main() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":9090", s.Router()))
}

