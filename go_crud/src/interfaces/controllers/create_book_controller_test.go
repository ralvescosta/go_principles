package controllers

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	stuntman "crud/src/interfaces/__test__"

	core "crud/src/__core__"
	errors "crud/src/__core__/errors"
	entities "crud/src/business/entities"
)

type mocksCreateStruct struct {
	inputCreateBook *entities.InputCreateBook
	bookEntity      *entities.BookEntity
	request         *core.HTTPRequest
	usecaseSpy      *stuntman.BookSpy
}

func makeCreateMocks() *mocksCreateStruct {

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

	request := &core.HTTPRequest{
		Body: inputCreateBook,
	}

	usecaseSpy := &stuntman.BookSpy{}

	return &mocksCreateStruct{
		inputCreateBook: inputCreateBook,
		bookEntity:      bookEntity,
		request:         request,
		usecaseSpy:      usecaseSpy,
	}
}

func TestCreateController(t *testing.T) {
	mocks := makeCreateMocks()
	mocks.usecaseSpy.On("RegisterABook", mocks.inputCreateBook).Return(mocks.bookEntity, nil)

	sut := CreateABookController(mocks.usecaseSpy)
	response := sut.Handle(mocks.request)

	assert.Equal(t, response.StatusCode, 201)
}

func TestCreateIfSomeErrorOccurOnUsecaseReturnsStatusCode500(t *testing.T) {
	mocks := makeCreateMocks()
	mocks.usecaseSpy.On("RegisterABook", mocks.inputCreateBook).Return(mocks.bookEntity, &errors.InternalServerError{})

	sut := CreateABookController(mocks.usecaseSpy)
	response := sut.Handle(mocks.request)

	assert.Equal(t, response.StatusCode, 500)
}
