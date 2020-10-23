package repositories

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"

	entities "crud/src/business/entities"
)

type mocksStruct struct {
	db              *sql.DB
	dbMocked        sqlmock.Sqlmock
	dbErr           error
	tableRows       []string
	inputCreateBook *entities.InputCreateBook
	bookEntity      *entities.BookEntity
}

func makeMocks() *mocksStruct {
	db, dbMocked, dbErr := sqlmock.New()

	tableRows := []string{"id", "title", "author", "publishing_company", "edition", "created_at", "updated_at", "deleted_at"}

	inputCreateBook := &entities.InputCreateBook{
		Title:             "title",
		Author:            "author",
		PublishingCompany: "publishing",
		Edition:           1,
	}

	bookEntity := &entities.BookEntity{
		ID:                1,
		Title:             "title",
		Author:            "author",
		PublishingCompany: "publishing",
		Edition:           1,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		DeletedAt:         sql.NullTime{},
	}

	return &mocksStruct{
		db:              db,
		dbMocked:        dbMocked,
		dbErr:           dbErr,
		tableRows:       tableRows,
		inputCreateBook: inputCreateBook,
		bookEntity:      bookEntity,
	}
}

func TestBooksRepositoryCreateMethod(t *testing.T) {
	mocks := makeMocks()

	dbMocked := mocks.dbMocked
	db := mocks.db
	input := mocks.inputCreateBook
	entity := mocks.bookEntity
	query := "INSERT INTO books \\(title, author, publishing_company, edition\\) VALUES \\(\\?, \\?, \\?, \\?\\) RETURNING *"
	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows(mocks.tableRows).
		AddRow(entity.ID, entity.Title, entity.Author, entity.PublishingCompany, entity.Edition, entity.CreatedAt, entity.UpdatedAt, entity.DeletedAt)

	dbMocked.ExpectPrepare(query).
		ExpectQuery().
		WithArgs(input.Title, input.Author, input.PublishingCompany, input.Edition).
		WillReturnRows(rows)

	sut := BooksRepository(db)

	_, err := sut.Create(input)

	assert.NoError(t, err)
}

func TestBooksRepositoryFindByIDMethod(t *testing.T) {
	mocks := makeMocks()

	dbMocked := mocks.dbMocked
	db := mocks.db
	entity := mocks.bookEntity
	query := "SELECT id, title, author, publishing_company, edition, created_at, updated_at, deleted_at FROM books WHERE id = \\?"
	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows(mocks.tableRows).
		AddRow(entity.ID, entity.Title, entity.Author, entity.PublishingCompany, entity.Edition, entity.CreatedAt, entity.UpdatedAt, entity.DeletedAt)

	dbMocked.
		ExpectPrepare(query).
		ExpectQuery().
		WithArgs(entity.ID).
		WillReturnRows(rows)

	sut := BooksRepository(db)

	result, err := sut.FindByID(entity.ID)

	assert.NotNil(t, result)
	assert.NoError(t, err)
}

func TestBooksRepositoryFindAllMethod(t *testing.T) {
	mocks := makeMocks()

	dbMocked := mocks.dbMocked
	db := mocks.db
	entity := mocks.bookEntity
	query := "SELECT id, title, author, publishing_company, edition, created_at, updated_at, deleted_at FROM books"
	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows(mocks.tableRows).
		AddRow(entity.ID, entity.Title, entity.Author, entity.PublishingCompany, entity.Edition, entity.CreatedAt, entity.UpdatedAt, entity.DeletedAt)

	dbMocked.
		ExpectQuery(query).
		WillReturnRows(rows)

	sut := BooksRepository(db)

	result, err := sut.FindAll()

	assert.NotEmpty(t, result)
	assert.NoError(t, err)
	assert.Len(t, *result, 1)
}

func TestBooksRepositoryUpdateMethod(t *testing.T) {
	mocks := makeMocks()

	dbMocked := mocks.dbMocked
	db := mocks.db
	entity := mocks.bookEntity
	input := mocks.inputCreateBook
	query := "UPDATE books SET author = 'author', edition = 1, publishing_company = 'publishing', title = 'title' WHERE id = \\? RETURNING *"
	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows(mocks.tableRows).
		AddRow(entity.ID, entity.Title, entity.Author, entity.PublishingCompany, entity.Edition, entity.CreatedAt, entity.UpdatedAt, entity.DeletedAt)

	dbMocked.
		ExpectPrepare(query).
		ExpectQuery().
		WithArgs(entity.ID).
		WillReturnRows(rows)

	sut := BooksRepository(db)

	result, err := sut.Update(entity.ID, input)

	assert.NotEmpty(t, result)
	assert.NoError(t, err)
}

func TestBooksRepositoryDeleteMethod(t *testing.T) {
	mocks := makeMocks()

	dbMocked := mocks.dbMocked
	db := mocks.db
	entity := mocks.bookEntity
	query := "DELETE FROM books WHERE id = \\? RETURNING *"
	defer func() {
		db.Close()
	}()

	rows := sqlmock.NewRows(mocks.tableRows).
		AddRow(entity.ID, entity.Title, entity.Author, entity.PublishingCompany, entity.Edition, entity.CreatedAt, entity.UpdatedAt, entity.DeletedAt)

	dbMocked.
		ExpectPrepare(query).
		ExpectQuery().
		WithArgs(entity.ID).
		WillReturnRows(rows)

	sut := BooksRepository(db)

	result, err := sut.Delete(entity.ID)

	assert.NotEmpty(t, result)
	assert.NoError(t, err)
}
