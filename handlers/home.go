package handlers

import (
	"net/http"
	"log"
	"html/template"

	"github/grovercoder/syndio/datastore"
)

/*
	This file is not required to meet the stated needs.
	It is used to help validate the results
*/

// HomeHandler handles the "/" route, serving the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	employees, err := datastore.GetEmployeeData()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		log.Printf("Database error: %v", err)
		return
	}

	// We can add the .Funcs(funcMap) code here if we ever need template helpers
	tmpl := template.New("index.html")
	tmpl, err = tmpl.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
		return
	}

	// Render the template to the response stream
	err = tmpl.Execute(w, employees)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		log.Printf("Template execution error: %v", err)
		return
	}
}