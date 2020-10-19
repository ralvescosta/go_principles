package main

import (
	database "crud/src/frameworks/database"
)

func main() {
	_dbConnection := database.DbConnection()
	dbConn, err := _dbConnection.Connect()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	dbConn.Exec(`
		CREATE TABLE books
		(
			id SERIAL NOT NULL PRIMARY KEY,
			title varchar(255) NOT NULL,
			author varchar(255) NOT NULL,
			publishing_company varchar(255) NOT NULL,
			edition int,
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP WITH TIME ZONE
		);
	`)
}
