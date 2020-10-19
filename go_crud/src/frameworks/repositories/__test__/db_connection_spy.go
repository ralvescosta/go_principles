package test

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

// DbConnectionSpy ...
type DbConnectionSpy struct {
	mock.Mock
}

// Connect ...
func (d *DbConnectionSpy) Connect() (*sql.DB, error) {
	args := d.Called()

	return args.Get(0).(*sql.DB), args.Error(1)
}
