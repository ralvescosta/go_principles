package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres driver
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "12345"
	DB_DATABASE = "default"
)

type dbConnection struct{}

// Connect ...
func (*dbConnection) Connect() (*sql.DB, error) {
	psql := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_DATABASE)

	db, err := sql.Open("postgres", psql)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// DbConnection ...
func DbConnection() IDbConnection {
	return &dbConnection{}
}
