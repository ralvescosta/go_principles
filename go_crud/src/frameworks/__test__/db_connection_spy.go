package test

import "database/sql"

type SQLDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type MockDB struct {
	callParams []interface{}
}

func (mdb *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	mdb.callParams = []interface{}{query}
	mdb.callParams = append(mdb.callParams, args...)

	return nil, nil
}

// Add a helper method to inspect the `callParams` field
func (mdb *MockDB) CalledWith() []interface{} {
	return mdb.callParams
}
