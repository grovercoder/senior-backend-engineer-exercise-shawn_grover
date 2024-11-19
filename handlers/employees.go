package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	
	"github/grovercoder/syndio/datastore"
	"github/grovercoder/syndio/models"
)

/*
	This would be the entry point for the Employees endpoint of an API.

	For the current stated needs, we simply handle the ingestion of JSON based data
*/

func JobDataIngestion(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming JSON body
    var jobDataArray []models.JobDataJSON
    err := json.NewDecoder(r.Body).Decode(&jobDataArray)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
        return
    }
	
	// use the array to update the database accordingly
	err = datastore.IngestJobData(jobDataArray)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ingest error: %v", err), http.StatusBadRequest)
		log.Printf("Ingest error: %v", err)
		return
	}
    
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
}
