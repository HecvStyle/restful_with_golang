package dbutils

import (
	"database/sql"
	"log"
)

// Initialize ...
func Initialize(dbDriver *sql.DB) {
	statement, driverError := dbDriver.Prepare(train)
	if driverError != nil {
		log.Println(driverError)
	}

	_, statementError := statement.Exec()
	if statementError != nil {
		log.Println("Table already exists!")
	}

	statement, _ = dbDriver.Prepare(station)
	statement.Exec()

	statement, _ = dbDriver.Prepare(schedule)
	statement.Exec()
	log.Println("All tables create/initialized successfully!")

}
