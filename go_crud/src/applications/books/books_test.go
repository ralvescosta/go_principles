package books

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	stuntman "crud/src/applications/__test__"
	entities "crud/src/business/entities"
)

type mockStruct struct {
	inputCreateBook *entities.InputCreateBook
	bookEntity      *entities.BookEntity
	booksEntity     *[]entities.BookEntity
	repositorySpy   *stuntman.BooksRepositorySpy
}

func makeMocks() *mockStruct {

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

	booksEntity := &[]entities.BookEntity{
		*bookEntity,
	}

	repositorySpy := &stuntman.BooksRepositorySpy{}

	return &mockStruct{
		inputCreateBook: inputCreateBook,
		bookEntity:      bookEntity,
		booksEntity:     booksEntity,
		repositorySpy:   repositorySpy,
	}
}

func TestRegisterABook(t *testing.T) {
	mocks := makeMocks()

	mocks.repositorySpy.On("Create", mocks.inputCreateBook).Return(mocks.bookEntity, nil)

	sut := Books(mocks.repositorySpy)

	result, err := sut.RegisterABook(mocks.inputCreateBook)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestFindABook(t *testing.T) {
	mocks := makeMocks()

	mocks.repositorySpy.On("FindByID", mocks.bookEntity.ID).Return(mocks.bookEntity, nil)

	sut := Books(mocks.repositorySpy)

	result, err := sut.FindABook(mocks.bookEntity.ID)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestGetAllBooks(t *testing.T) {
	mocks := makeMocks()

	mocks.repositorySpy.On("FindAll").Return(mocks.booksEntity, nil)

	sut := Books(mocks.repositorySpy)

	result, err := sut.GetAllBooks()

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestUpdateABook(t *testing.T) {
	mocks := makeMocks()

	mocks.repositorySpy.On("Update", mocks.bookEntity.ID, mocks.inputCreateBook).Return(mocks.bookEntity, nil)

	sut := Books(mocks.repositorySpy)

	result, err := sut.UpdateABook(mocks.bookEntity.ID, mocks.inputCreateBook)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestDeleteABook(t *testing.T) {
	mocks := makeMocks()

	mocks.repositorySpy.On("Delete", mocks.bookEntity.ID).Return(mocks.bookEntity, nil)

	sut := Books(mocks.repositorySpy)

	result, err := sut.DeleteABook(mocks.bookEntity.ID)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}
