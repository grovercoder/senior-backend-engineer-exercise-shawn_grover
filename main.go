package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github/grovercoder/syndio/handlers" 
	"github/grovercoder/syndio/datastore" 
)

func main() {
	// ensure the job_data table is created 
	// - other operations will depend on this table and it is best to check this early to avoid redundant checks
	datastore.EnsureJobDataTable()
	
	// Retrieve the port using the getPort function
	port := getPort()

	// assign the route handlers
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/api/ingest", handlers.JobDataIngestion)

	// Start the web server
	log.Printf("Starting server on port %d...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getPort() int {
	portStr := os.Getenv("WEB_PORT")
	if portStr == "" {
		portStr = "51234" // Default port as a string
	}

	// Convert portStr to an integer
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	return port
}