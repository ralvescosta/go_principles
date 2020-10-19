package database

import (
	"database/sql"
)

// IDbConnection ...
type IDbConnection interface {
	Connect() (*sql.DB, error)
}
