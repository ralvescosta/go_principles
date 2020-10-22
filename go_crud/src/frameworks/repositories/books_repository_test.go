package repositories

import (
	"database/sql"
	"fmt"
	"testing"

	entities "crud/src/business/entities"
)

type MockDB struct {
}



func (mdb *MockDB) Prepare(query string) (*sql.Stmt, error) {
	return nil, nil
}
func (mdb *MockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func TestBooksRepository(t *testing.T) {

	sut := BooksRepository(&MockDB{})

	result, err := sut.Create(&entities.InputCreateBook{Author: "Author", Edition: 1, PublishingCompany: "PublishingCompany", Title: "Title"})

	fmt.Println(result, err)
	// assert.Equal(t, true, true, "SigninUsecase()")

}
