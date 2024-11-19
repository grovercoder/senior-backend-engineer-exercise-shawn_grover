package datastore

import (
	"database/sql"
	"fmt" 

	"github/grovercoder/syndio/models"
	_ "github.com/mattn/go-sqlite3"

)

/*
	This file provides the database access and manipulation logic.

	This would normally be organized into a better structure, 
	but this is sufficient for the current needs.
*/

// Where the SQLITE database resides
const DBPATH = "./data/employees.db"


// Create the job_data table if needed
func EnsureJobDataTable() error {
	// Connect to the database
	db, err := sql.Open("sqlite3", DBPATH)
	if err != nil {
		return fmt.Errorf("database open error: %w", err)
	}
	defer db.Close()

	// SQL to create the "job_data" table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS job_data (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		employee_id INTEGER,
		department TEXT NOT NULL,
		job_title TEXT NOT NULL
	);`

	// Execute the statement
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("table creation error: %w", err)
	}

	// log.Println("Verified or created the job_data table")
	return nil
}

// retreive employee data, combining the employees and job_data tables
func GetEmployeeData() ([]models.EmployeeData, error) {
	// Connect to the database
	db, err := sql.Open("sqlite3", DBPATH)
	if err != nil {
		return nil, fmt.Errorf("database open error: %w", err)
	}
	defer db.Close()

	// Retrieve the employee records
	rows, err := db.Query("SELECT e.id, e.gender, j.department, j.job_title from employees e, job_data j where j.employee_id = e.id")
	if err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}
	defer rows.Close()

	// Create an array of EmployeeRecord objects
	var EmployeeInfo []models.EmployeeData
	for rows.Next() {
		var row models.EmployeeData
		if err := rows.Scan(&row.ID, &row.Gender, &row.Department, &row.JobTitle); err != nil {
			return nil, fmt.Errorf("database scan error: %w", err)
		}
		EmployeeInfo = append(EmployeeInfo, row)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	// Return the retreived data
	return EmployeeInfo, nil
}

// Handle ingesting job data from an array.
// NOTE: An assumption has been made - An employee can only have one job defined at a time.
//       Based on this assumption, we try to update an existing record. If no rows were modified
// 		 an insert statement is then done.  This logic may not be correct and will be subject to 
//       change pending more information.
func IngestJobData(jobDataArray []models.JobDataJSON) error {
	err := EnsureJobDataTable()
	if err != nil {
        return fmt.Errorf("Error ensuring job data table: %v", err)
    }

    // Open the database connection
    db, err := sql.Open("sqlite3", DBPATH)
    if err != nil {
        return fmt.Errorf("Database error: %v", err)
    }
    defer db.Close()

    // Iterate over the job data and insert or update job_data table
    for _, job := range jobDataArray {
        // Check if the employee_id exists in the employees table
        var exists bool
        err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM employees WHERE id = ?)", job.EmployeeID).Scan(&exists)
        if err != nil || !exists {
            return fmt.Errorf("Employee with ID %d does not exist", job.EmployeeID)
        }

        // Attempt to update the job data for the employee
        result, err := db.Exec(`
            UPDATE job_data 
            SET department = ?, job_title = ?
            WHERE employee_id = ?`, job.Department, job.JobTitle, job.EmployeeID)

        if err != nil {
            return fmt.Errorf("Error updating job data: %v", err)
        }

        // If no rows were updated, insert a new record
        rowsAffected, err := result.RowsAffected()
        if err != nil {
            return fmt.Errorf("Error checking rows affected: %v", err)
        }

        if rowsAffected == 0 {
            // Insert the job data if no existing record was updated
            _, err = db.Exec(`
                INSERT INTO job_data (employee_id, department, job_title)
                VALUES (?, ?, ?)`, job.EmployeeID, job.Department, job.JobTitle)
            if err != nil {
                return fmt.Errorf("Error inserting job data: %v", err)
            }
        }
    }

    return nil
}
