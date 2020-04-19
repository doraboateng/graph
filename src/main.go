package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kwcay/boateng-graph-service/src/router"
)

// Build-time variables.
var version string
var gitHash string

func main() {
	router := router.Create()
	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "80"
	}

	// Launch server
	log.Println("Environment: " + os.Getenv("APP_ENV"))
	log.Println("Version: " + version)
	log.Println("Hash: " + gitHash)
	log.Printf("Serving GraphQL at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
