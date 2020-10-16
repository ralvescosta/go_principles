package test

import (
	database "study/src/database"

	"github.com/stretchr/testify/mock"
)

// DatabaseSpy ...
type DatabaseSpy struct {
	mock.Mock
	database.IDatabase
}

// GetTable ...
func (d *DatabaseSpy) GetTable() string {
	args := d.Called()
	return args.String(0)
}
