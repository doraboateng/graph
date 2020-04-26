package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kwcay/boateng-graph-service/src/dgraph"
	"github.com/kwcay/boateng-graph-service/src/router"
)

// Build-time variables
var version string
var gitHash string

func main() {
	// Setup router
	router := router.Create()
	port := os.Getenv("API_PORT")

	if port == "" {
		port = "80"
	}

	// Load Graph schema
	dgraph.RefreshSchema()

	// Launch server
	log.Println("Environment: " + os.Getenv("API_ENV"))
	log.Println("Version: " + version)
	log.Println("Hash: " + gitHash)
	log.Printf("Serving GraphQL at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
