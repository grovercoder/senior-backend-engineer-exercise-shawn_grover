package models

// Matching the employees table
type EmployeeRecord struct {
	ID    	int     `db:"id"`
	Gender 	string  `db:"gender"`
}

// Matching the job_data table
type JobDataRecord struct {
	ID			int		`db:"id"`
	EmployeeID  int		`db:"employee_id"`
	Department	string	`db:"department"`
	JobTitle	string	`db:"job_title"`
}

// Used with a query combining employees and job_data records
type EmployeeData struct {
	ID    	int     `db:"id"`
	Gender 	string  `db:"gender"`
	Department	string	`db:"department"`
	JobTitle	string	`db:"job_title"`	
}	

// For handling incoming JSON data
type JobDataJSON struct {
	EmployeeID  int		`json:"employee_id"`
	Department	string	`json:"department"`
	JobTitle	string	`json:"job_title"`
}
