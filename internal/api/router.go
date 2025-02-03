package api

import (
	"log"
	"net/http"
	"os"
)

func NewServer() *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	return &http.Server{Addr: "localhost:" + port}
}
