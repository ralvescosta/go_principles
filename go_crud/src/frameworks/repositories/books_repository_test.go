package repositories

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	entities "crud/src/business/entities"
)

func Prepare(query string) (*sql.Stmt, error) {
	return &sql.Stmt{}, nil
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return &sql.Rows{}, nil
}

func TestBooksRepository(t *testing.T) {
	db, mock, _ := sqlmock.New()

	dbConn := &DatabaseStruct{
		Prepare: db.Prepare,
		Query:   db.Query,
	}

	mock.ExpectPrepare("INSERT INTO books (title, author, publishing_company, edition) VALUES ($1, $2, $3, $4) RETURNING *").ExpectQuery()

	sut := BooksRepository(dbConn)

	result, err := sut.Create(&entities.InputCreateBook{Author: "Author", Edition: 1, PublishingCompany: "PublishingCompany", Title: "Title"})

	fmt.Println(result, err)
	// assert.Equal(t, true, true, "SigninUsecase()")

}
