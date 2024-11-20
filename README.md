# Syndio Backend App Submission

See the [original README](https://github.com/syndio/senior-backend-engineer-exercise/blob/main/README.md) for the original task.

## Problem Definition

An SQLite3 database is provided that contains an `employees` table.

The task is to handle ingesting "job data" for employees.

## Quick Analysis

1. A `job_data` table will need to be created
1. An API is required that will allow the job_data table to be populated from a JSON array. This json array must contain objects that have "employee_id", "department", and "job_title" properties
1. The PORT for the API must be defined by an environment variable.
1. The solution must be written in Golang.

## Usage

1. Clone/Fork/Copy this project to a working directory
1. Open a command prompt and change into the project directory.


```bash
cd project_directory

# retreive / reset the database
chmod +x reset_db.sh
./reset_db.sh

# run the API service on the default port (51234)
go run main.go

# run the API on a different port
# (replace the '55555' value with your desired port)
WEB_PORT=55555 go run main.go

# Alternatively, build an executable:
go build -o myapp main.go

# And run the executable
./myapp
# or
WEB_PORT=55555 ./myapp

```

Now that the service is running, you can ingest the job data by calling the `http://localhost:PORT/api/ingest`, with the JSON object passed as the body of the request.  You can do this with a `curl` command, postman, or other REST API client:

```bash
curl -X POST http://localhost:51111/api/ingest -d '[
  { "employee_id": 1, "department": "Engineering", "job_title": "Senior Engineer" },
  { "employee_id": 2, "department": "Engineering", "job_title": "Super Senior Engineer" },
  { "employee_id": 3, "department": "Sales", "job_title": "Head of Sales" },
  { "employee_id": 4, "department": "Support", "job_title": "Tech Support" },
  { "employee_id": 5, "department": "Engineering", "job_title": "Junior Engineer" },
  { "employee_id": 6, "department": "Sales", "job_title": "Sales Rep" },
  { "employee_id": 7, "department": "Marketing", "job_title": "Senior Marketer" }
]'
```

You can inspect the database to ensure the job_data table has been created and populated accordingly.  You can also load `http://localhost:PORT/` in a browser to see the data from the employees and job_data tables joined.  (loading this before ingesting would show no data, or possibly result in an error)

## Notes:

1. **ASSUMPTION** - An assumption has been made that an employee may have zero or only one job data entry.  This simplified the logic and avoided dealing with multiple job_data records for an employee.  It is understood this assumption may be incorrect and extra effort would be needed if a zero or many relationship is required.
1. This code is not production ready.  It is not currently configured to handle penetration testing, authentication, or other common aspects.  There may be error conditions that are not properly caught/handled.
1. The ingestion endpoint would normally be a POST operation.  I intentionally used a GET operation to simplify development and testing.
1. Authorization should always be validated when modifying data.  This is outside the scope of this introduction project though and has been omitted.  I suspect this is an understood omission to the requirements.
1. I have applied my systems knowledge from other languages/frameworks here in terms of how I organized the code.  I feel this organization allows for easy maintenance while allowing for growth over time.  It is possible to provide all the capabilities in a single file, but this would be a nightmare to maintain or grow.
1. Coding the solution took me approx 2 hours.  However, this can be attributed to my relative new-ness to Go - the task itself and the concepts involved are straight forward for me.  I also spent almost as much time documenting and vetting the process as I did writing code.


